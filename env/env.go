package env

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func ENV(key string) string {

	err := godotenv.Load("env/.env")

	if err != nil {
		log.Fatal("ENV NOT OPEN")

	}
	return os.Getenv(key)

}
