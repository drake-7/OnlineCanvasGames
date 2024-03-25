package env

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func GetEnvVal(key string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	log.Println("Couldn't find value in environment:", key)
	return ""
}

func InitEnv() {
	if err := godotenv.Load(); err != nil {
		log.Println("Couldn't find .env file")
	}
}
