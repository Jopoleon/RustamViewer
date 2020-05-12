package storage

import (
	"github.com/Jopoleon/rustamViewer/config"
	"github.com/Jopoleon/rustamViewer/logger"
	"github.com/Jopoleon/rustamViewer/storage/ftp"
	"github.com/Jopoleon/rustamViewer/storage/postgres"
	"github.com/pkg/errors"
)

type Storage struct {
	Logger *logger.LocalLogger
	DB     *postgres.DB
	FTP    *ftp.LocalFTP
}

func NewStorage(cfg config.Config, log *logger.LocalLogger) (*Storage, error) {
	res := &Storage{}
	db, err := postgres.NewPostgres(cfg, log)
	if err != nil {
		return res, errors.WithStack(err)
	}

	ftps, err2 := ftp.NewFTP(&cfg, log)
	if err2 != nil {
		return res, errors.WithStack(err2)
	}
	res.FTP = ftps
	res.DB = db
	res.Logger = log
	return res, nil
}
