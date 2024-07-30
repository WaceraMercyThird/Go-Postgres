package initializer

import (
	"Go_conn/entities"
	"github.com/sirupsen/logrus"
)

func AutoMigrate() {
	PostgresDB.AutoMigrate(
		&entities.ReportingTransaction{})
	logrus.Info("Database migrations successful")
}
