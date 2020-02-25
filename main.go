package main

import (
	"github.com/Jopoleon/rustamViewer/config"
	"github.com/sirupsen/logrus"

	"github.com/Jopoleon/rustamViewer/app"
)

func main() {
	ll := logrus.New()
	ll.SetReportCaller(true)
	cfg := config.NewConfig()

	a, err := app.New(cfg, ll)
	if err != nil {
		logrus.Fatal(err)
	}
	a.Run()

}
