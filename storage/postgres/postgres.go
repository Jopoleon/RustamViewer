package postgres

import (
	"fmt"

	"github.com/Jopoleon/rustamViewer/logger"

	"github.com/Jopoleon/rustamViewer/config"
	_ "github.com/lib/pq"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

type DB struct {
	Logger   *logger.LocalLogger
	DB       *sqlx.DB
	DBConfig *config.Config
}

func NewPostgres(cfg config.Config, logger *logger.LocalLogger) (*DB, error) {
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

	res := &DB{DB: db, Logger: logger, DBConfig: &cfg}
	//err = res.Migrate()
	//if err != nil {
	//	logger.Errorf("could not establish connection to ", ommitesStr, err)
	//	return nil, errors.WithStack(err)
	//}
	return res, nil
}
