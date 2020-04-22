package postgres

import (
	"fmt"

	"github.com/Jopoleon/rustamViewer/config"
	_ "github.com/lib/pq"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type DB struct {
	Logger   *logrus.Logger
	DB       *sqlx.DB
	DBConfig *config.Config
}

func NewPostgres(cfg config.Config, logger *logrus.Logger) (*DB, error) {
	str := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPass, cfg.DBName)
	ommitesStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, "[ommited]", cfg.DBName)
	db, err := sqlx.Connect("postgres", str)
	if err != nil {
		logger.Errorf("could not establish connection to ", ommitesStr, err)
		return nil, errors.WithStack(err)
	}
	db.SetMaxOpenConns(5)
	db.SetMaxIdleConns(5)
	return &DB{DB: db, Logger: logger, DBConfig: &cfg}, nil
}
