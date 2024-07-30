package main

import "Go_conn/initializer"

func main() {

	// Log file name
	var configFile = ".env"

	var logFile = "/Go_logs.log"

	initializer.LoadEnv(configFile)
	initializer.ConnectToPostgresDB()
	initializer.Logger(logFile)
	initializer.AutoMigrate()

}
