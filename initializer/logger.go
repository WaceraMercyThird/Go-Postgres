package initializer

import (
	"Go_conn/entities"
	"github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
	"log"
	"time"
)

// InitLogger initializes the logger
func Logger(logFile string) {
	config := entities.ConfigDb{}
	logg := config.FileSetting.SystemLogs
	// Create log file with rotation
	logFilePath := logg + logFile + ".%Y%m%d"

	logFileWriter, err := rotatelogs.New(

		logFilePath,
		rotatelogs.WithMaxAge(time.Hour*24),    // Rotate daily
		rotatelogs.WithRotationTime(time.Hour), // Rotate every hour
	)
	if err != nil {

		logrus.Fatalf("Failed to create log file: %v", err)
	}

	log.SetOutput(logFileWriter)

	logrus.SetOutput(logFileWriter)

	logrus.SetFormatter(&logrus.JSONFormatter{})

	logrus.SetLevel(logrus.InfoLevel)
}
