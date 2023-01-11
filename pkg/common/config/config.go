package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	PORT 				string
	HOST 				string
	DB_USER 		string
	DB_PASSWORD string
	DB_NAME 		string
	JWT_TOKEN		string
}

func InitConfig() (Config, error) {
	var c Config
	err := godotenv.Load()

	if err != nil {
		log.Fatalln(err)
	}

	c.PORT = os.Getenv("PORT")
	c.HOST = os.Getenv("HOST")
	c.DB_USER = os.Getenv("DB_USER")
	c.DB_PASSWORD = os.Getenv("DB_PASSWORD")
	c.DB_NAME = os.Getenv("DB_NAME")
	c.JWT_TOKEN = os.Getenv("JWT_TOKEN")

	return c, err
}