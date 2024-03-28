package env

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func GetEnvVal(key string) string {
	value, ok := os.LookupEnv(key)
	if !ok {
		log.Println("couldn't find environment value:", key)
	}
	return value
}

func InitEnv() {
	if err := godotenv.Load(); err != nil {
		log.Println("Couldn't find .env file")
	}
}
