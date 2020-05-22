package app

import (
	"time"

	"github.com/Jopoleon/rustamViewer/logger"

	"github.com/Jopoleon/rustamViewer/api"
	"github.com/Jopoleon/rustamViewer/config"
	"github.com/Jopoleon/rustamViewer/storage"
)

// App struct is base struct with all essential information about application
type App struct {
	StartTime time.Time
	Logger    *logger.LocalLogger
	Config    *config.Config
}

// New inits new App instance
func New(cfg *config.Config, l *logger.LocalLogger) (*App, error) {
	return &App{
		Logger:    l,
		Config:    cfg,
		StartTime: time.Now(),
	}, nil
}

func (a *App) Run() {
	st, err := storage.NewStorage(*a.Config, a.Logger)
	if err != nil {
		a.Logger.Fatalln("can't create new storage: ", err)
	}
	appi := api.NewAPI(st, a.Logger, a.Config)
	appi.InitRouter()
	api.ServeAPI(appi)
}
