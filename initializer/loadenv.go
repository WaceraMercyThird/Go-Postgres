package initializer

import (
	"fmt"
	"github.com/joho/godotenv"
)

func LoadEnv(configFile string) {
	err := godotenv.Load(configFile)
	if err != nil {
		fmt.Print("Error while loading the env")
	}
	fmt.Println("Loaded env successfully")
}
