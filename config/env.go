package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Env struct {
	MongoURL     string
	DatabaseName string
}

var Envs Env

func InitializeEnv() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	Envs = Env{
		MongoURL:     os.Getenv("MONGO_URL"),
		DatabaseName: os.Getenv("DB_NAME"),
	}
}
