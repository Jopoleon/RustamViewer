package controllers

import (
	"net/http"
	"time"

	"github.com/Jopoleon/rustamViewer/models"

	"github.com/Jopoleon/rustamViewer/config"
	"github.com/Jopoleon/rustamViewer/storage"
	"github.com/go-chi/chi"
	"github.com/sirupsen/logrus"
)

const (
	CookieMaxAge   = 259200 //cookie expiration time in seconds (3 days)
	CookieAuthName = "AUTH_SESSION"
	CookieName     = "user_session"
)

type Controllers struct {
	StartTime  time.Time
	Logger     *logrus.Logger
	HttpPort   string
	Config     *config.Config
	Router     *chi.Mux
	Repository *storage.Storage
}

func NewControllers(rep *storage.Storage, log *logrus.Logger, cfg *config.Config) *Controllers {
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
