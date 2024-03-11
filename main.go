package main

import (
	"os"

	"github.com/Marvit-Solutions/csw-golang/library/config"

	"github.com/sirupsen/logrus"
)

// init sets the timezone environment variable to Asia/Jakarta.
func init() {
	_ = os.Setenv("TZ", "Asia/Jakarta")
}

func main() {
	// Configure logrus formatter.
	logrus.SetFormatter(&logrus.TextFormatter{
		ForceColors:     true,
		FullTimestamp:   true,
		TimestampFormat: "02-01-2024 15:04:05",
	})

	// Load environment variables from config.
	env = config.NewConfig()

	logrus.Info("success loading .env")
	logrus.Info("Starting server... \n")

	// Start the application.
	startApp()
}
