package main

import (
	"runtime"

	"github.com/Jopoleon/rustamViewer/app"
	"github.com/Jopoleon/rustamViewer/config"
	"github.com/Jopoleon/rustamViewer/logger"
	"github.com/sirupsen/logrus"
)

var isLocalRun bool

func main() {
	if runtime.GOOS == "windows" {
		isLocalRun = true
	}
	ll := logger.NewLogger(isLocalRun)
	cfg := config.NewConfig()

	a, err := app.New(cfg, ll)
	if err != nil {
		logrus.Fatal(err)
	}
	a.Run()
}
