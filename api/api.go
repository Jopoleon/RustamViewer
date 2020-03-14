package api

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Jopoleon/rustamViewer/api/controllers"

	"github.com/Jopoleon/rustamViewer/config"

	"github.com/Jopoleon/rustamViewer/storage"
	"github.com/go-chi/chi"
	"github.com/sirupsen/logrus"
)

type API struct {
	*controllers.Controllers
	StartTime  time.Time
	Logger     *logrus.Logger
	HttpPort   string
	Config     *config.Config
	Router     *chi.Mux
	Repository *storage.Storage
}

func NewAPI(rep *storage.Storage, log *logrus.Logger, cfg *config.Config) *API {
	a := &API{
		controllers.NewControllers(rep, log, cfg),
		time.Now(),
		log,
		cfg.HttpPort,
		cfg,
		chi.NewMux(),
		rep,
	}
	a.InitRouter()
	return a
}
func ServeAPI(api *API) {

	s := &http.Server{
		Addr:        "127.0.0.1:" + api.HttpPort,
		Handler:     api.Router,
		ReadTimeout: 1 * time.Minute,
	}
	//implementing graceful shutdown due to kubernetes sigterm
	idleConnsClosed := make(chan struct{})
	go func() {
		sigint := make(chan os.Signal, 1)

		// interrupt signal sent from terminal
		signal.Notify(sigint, os.Interrupt)
		// sigterm signal sent from kubernetes, Kubernetes sends a SIGTERM signal which is different from SIGINT (Ctrl+Client).
		signal.Notify(sigint, syscall.SIGTERM)

		<-sigint
		// We received an interrupt signal, shut down.
		if err := s.Shutdown(context.Background()); err != nil {
			api.Logger.Errorf("HTTP server Shutdown: %+v \n", err)
		}
		//NIT: Каналы можно оставлять открытыми
		close(idleConnsClosed)
	}()
	api.Logger.Infof("serving api at http://%s", s.Addr)
	if err := s.ListenAndServe(); err != http.ErrServerClosed {
		api.Logger.Error(err)
		close(idleConnsClosed)
	}

	<-idleConnsClosed

}
