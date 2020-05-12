package ftp

import (
	"io/ioutil"

	"github.com/Jopoleon/rustamViewer/logger"

	"github.com/Jopoleon/rustamViewer/config"
	"github.com/pkg/errors"
)

type LocalFTP struct {
	Logger    *logger.LocalLogger
	FTPConfig *config.FTP
}

func NewFTP(cfg *config.Config, logger *logger.LocalLogger) (*LocalFTP, error) {
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
