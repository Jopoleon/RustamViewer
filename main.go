package main

import (
	"fmt"

	"github.com/Jopoleon/rustamViewer/app"
	"github.com/Jopoleon/rustamViewer/config"
	"github.com/Jopoleon/rustamViewer/logger"
	"github.com/sirupsen/logrus"
)

func main() {
	fmt.Println("starting......")
	cfg := config.NewConfig()

	ll := logger.NewLogger(cfg.ProductionStart)

	a, err := app.New(cfg, ll)
	if err != nil {
		logrus.Fatal(err)
	}
	a.Run()
}
