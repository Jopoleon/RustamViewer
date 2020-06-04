package main

import (
	"github.com/Jopoleon/rustamViewer/app"
	"github.com/Jopoleon/rustamViewer/config"
	"github.com/Jopoleon/rustamViewer/logger"
	"github.com/sirupsen/logrus"
)

func main() {

	cfg := config.NewConfig()

	ll := logger.NewLogger(cfg.ProductionStart)
	ll.Println("starting......")
	ll.Println("hello123")
	a, err := app.New(cfg, ll)
	if err != nil {
		logrus.Fatal(err)
	}
	a.Run()
}
