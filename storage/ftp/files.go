package ftp

import (
	"fmt"
	"os"

	"github.com/pkg/errors"
)

const PATH_FORMAT = "%s%s.%s"

func (ftp *LocalFTP) GetFile(fileName, fileType string) (*os.File, error) {
	path := fmt.Sprintf(PATH_FORMAT, ftp.FTPConfig.FilesPath, fileName, fileType)
	file, err := os.OpenFile(path, os.O_RDONLY, 0777)
	if err != nil {
		ftp.Logger.Errorf("%v", errors.WithStack(err))
		return nil, errors.WithStack(err)
	}
	return file, nil
}
