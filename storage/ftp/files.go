package ftp

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/pkg/errors"
)

const PATH_FORMAT = "%s%s.%s"

func (ftp *LocalFTP) GetFile(fileName, fileType string) (*os.File, error) {
	path := fmt.Sprintf(PATH_FORMAT, ftp.FTPConfig.FilesPath, fileName, fileType)
	//pp.Println(path)
	//list, err := ftp.ListFilesCallIDs()
	//if err != nil {
	//	ftp.Logger.Errorf("%v", errors.WithStack(err))
	//	return nil, errors.WithStack(err)
	//}
	//pp.Println(list)
	file, err := os.OpenFile(path, os.O_RDONLY, 0777)
	if err != nil {
		ftp.Logger.Errorf("%v", errors.WithStack(err))
		return nil, errors.WithStack(err)
	}
	return file, nil
}

func (ftp *LocalFTP) ListFilesCallIDs() ([]string, error) {
	var res []string
	files, err := ioutil.ReadDir(ftp.FTPConfig.FilesPath)
	if err != nil {
		ftp.Logger.Errorf("%v", errors.WithStack(err))
		return nil, errors.WithStack(err)
	}
	for _, f := range files {
		sl := strings.Split(f.Name(), "_")
		callID := sl[len(sl)-1]
		res = append(res, callID)
	}

	return res, nil
}
