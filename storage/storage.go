package storage

import (
	"github.com/Jopoleon/rustamViewer/config"
	"github.com/Jopoleon/rustamViewer/storage/ftp"
	"github.com/Jopoleon/rustamViewer/storage/postgres"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type Storage struct {
	Logger *logrus.Logger
	DB     *postgres.DB
	FTP    *ftp.LocalFTP
}

func NewStorage(cfg config.Config, logger *logrus.Logger) (*Storage, error) {
	res := &Storage{}
	db, err := postgres.NewPostgres(cfg, logger)
	if err != nil {
		return res, errors.WithStack(err)
	}

	ftps, err2 := ftp.NewFTP(&cfg, logger)
	if err2 != nil {
		return res, errors.WithStack(err2)
	}
	res.FTP = ftps
	res.DB = db
	res.Logger = logger
	return res, nil
}
