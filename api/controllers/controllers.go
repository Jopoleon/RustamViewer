package controllers

import (
	"net/http"
	"time"

	"github.com/Jopoleon/rustamViewer/logger"
	"github.com/Jopoleon/rustamViewer/models"

	"github.com/Jopoleon/rustamViewer/config"
	"github.com/Jopoleon/rustamViewer/storage"
	"github.com/go-chi/chi"
)

const (
	CookieMaxAge   = 259200 //cookie expiration time in seconds ( 259200 = 3 days )
	CookieAuthName = "AUTH_SESSION"
	CookieName     = "user_session"
)

type Controllers struct {
	StartTime  time.Time
	Logger     *logger.LocalLogger
	HttpPort   string
	Config     *config.Config
	Router     *chi.Mux
	Repository *storage.Storage
}

func NewControllers(rep *storage.Storage, log *logger.LocalLogger, cfg *config.Config) *Controllers {
	a := &Controllers{
		HttpPort:   cfg.HttpPort,
		StartTime:  time.Now(),
		Config:     cfg,
		Logger:     log,
		Repository: rep,
	}
	return a
}

func (c *Controllers) UserFromContext(w http.ResponseWriter, r *http.Request) *models.User {
	user, ok := r.Context().Value("userData").(models.User)
	if !ok || user.Login == "" {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return nil
	}
	return &user
}
