package logger

import (
	"fmt"
	"io"
	"os"
	"sync"
	"time"

	"github.com/pkg/errors"

	runtime "github.com/banzaicloud/logrus-runtime-formatter"
	"github.com/sirupsen/logrus"
)

type LocalLogger struct {
	*logrus.Logger
	Log         *logrus.Logger
	today       time.Time
	logFile     *os.File
	logFilePath string
	*sync.RWMutex
}

func NewLogger(productionStart string) *LocalLogger {

	switch productionStart {
	case "0":
		fmt.Println("logging to Stdout")
		return localLogger()
	case "1":
		fmt.Println("logging to local files in /projectRoot/logs and to Stdout")
		return productionLogger()
	default:
		fmt.Println("logging to Stdout")
		return localLogger()
	}
}

func localLogger() *LocalLogger {
	ll := logrus.New()
	childFormatter := logrus.TextFormatter{}
	runtimeFormatter := &runtime.Formatter{ChildFormatter: &childFormatter}
	runtimeFormatter.Line = true
	runtimeFormatter.File = true

	ll.SetFormatter(runtimeFormatter)
	res := LocalLogger{
		ll,
		ll,
		time.Now(),
		nil,
		"",
		new(sync.RWMutex),
	}
	return &res
}

func productionLogger() *LocalLogger {
	ll := logrus.New()
	childFormatter := logrus.JSONFormatter{}
	runtimeFormatter := &runtime.Formatter{ChildFormatter: &childFormatter}
	runtimeFormatter.Line = true
	runtimeFormatter.File = true

	ll.SetFormatter(runtimeFormatter)

	res := LocalLogger{
		ll,
		ll,
		time.Now(),
		nil,
		"",
		new(sync.RWMutex),
	}
	err := res.createLogFile()
	if err != nil {
		ll.Fatal(errors.WithStack(err))
		return nil
	}

	res.logFileWatcher()
	return &res
}
func (ll *LocalLogger) createLogFile() error {
	today := fmt.Sprintf("%d_%s_%d.txt",
		time.Now().Year(),
		time.Now().Month().String(),
		time.Now().Day())
	path := "./logs/" + today
	file, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		ll.Error(errors.WithStack(err))
		return errors.WithStack(err)
	}
	ll.Lock()
	ll.logFile = file
	ll.logFilePath = path
	mw := io.MultiWriter(os.Stdout, ll.logFile)
	ll.SetOutput(mw)
	ll.Unlock()
	return nil
}

func (ll *LocalLogger) logFileWatcher() {
	go func() {
		for {
			time.Sleep(1 * time.Hour)
			if ll.today.Second() != time.Now().Second() {
				err := ll.createLogFile()
				if err != nil {
					ll.Error(errors.WithStack(err))
				}
			}
		}
	}()
}
