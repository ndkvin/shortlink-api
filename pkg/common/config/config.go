package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func GetEnv(key string) (data string) {
	if err := godotenv.Load(); err != nil {
		log.Fatalln(err)
	}

	data = os.Getenv(key)
  return
}