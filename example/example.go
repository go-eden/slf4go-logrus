package main

import (
	slog "github.com/go-eden/slf4go"
	sl "github.com/go-eden/slf4go-logrus"
	"github.com/sirupsen/logrus"
	"os"
)

func init() {
	sl.Init()

	logrus.SetOutput(os.Stdout)
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetLevel(logrus.InfoLevel)
}

func main() {
	slog.Debug("hello")
	slog.Info("what???")
	slog.Warnf("warnning: %v", "surrender")

	log := slog.GetLogger()
	log.BindFields(slog.Fields{"logger": "test"})
	log.Errorf("error!!! %v", 100)
}
