package initializers

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func LoadEnvVariables() {

	env := os.Getenv("APP_ENV")
	if "" == env {
		env = "development"
	}
	err := godotenv.Load(".env." + env)

	if "test" == env {
		err = godotenv.Load(".env.test")
	}

	if err != nil {
		log.Fatal("Error loading .env file")
		return
	}

}
