package main

import (
	"csw-golang/library/config"
	"os"

	"github.com/sirupsen/logrus"
)

func init() {
	//godotenv.Load()
	_ = os.Setenv("TZ", "Asia/Jakarta")
}

func main() {
	logrus.SetFormatter(&logrus.TextFormatter{
		ForceColors:     true,
		FullTimestamp:   true,
		TimestampFormat: "02-01-2006 15:04:05",
	})

	env = config.NewConfig()

	logrus.Info("success loading .env")
	logrus.Info("Starting server... \n")

	startApp()
}
