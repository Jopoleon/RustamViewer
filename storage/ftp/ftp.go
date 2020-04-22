package ftp

import (
	"io/ioutil"

	"github.com/Jopoleon/rustamViewer/config"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type LocalFTP struct {
	Logger    *logrus.Logger
	FTPConfig *config.FTP
}

func NewFTP(cfg *config.Config, logger *logrus.Logger) (*LocalFTP, error) {
	ftp := &LocalFTP{}
	ftp.FTPConfig = &cfg.FTP
	ftp.Logger = logger
	files, err := ioutil.ReadDir(ftp.FTPConfig.FilesPath)
	if err != nil {
		logger.Errorf("could not read directory with wav and txt files, error: ", err)
		return nil, errors.WithStack(err)
	}
	if len(files) == 0 {
		logger.Warn("no wav and tt files in directory: ", ftp.FTPConfig.FilesPath)
	}
	return ftp, nil
}
