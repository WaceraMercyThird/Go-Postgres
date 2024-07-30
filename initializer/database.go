package initializer

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
)

var (
	PostgresDB *gorm.DB
)

func ConnectToPostgresDB() {
	log.Println("Connecting to Postgres Database....")
	var err error
	dsn := os.Getenv("PostgresDSN")
	PostgresDB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		log.Println("Failed to connect to postgres database")
		fmt.Println(err)
	} else {
		log.Println("Connected to postgres database")
	}
}
