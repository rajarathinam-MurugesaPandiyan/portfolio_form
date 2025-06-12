package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Env struct {
	MongoURL    string
	DatbaseName string
}

var Envs Env

func InitializeEnv() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	Envs = Env{
		MongoURL:    os.Getenv("MONGO_URL"),
		DatbaseName: os.Getenv("DB_NAME"),
	}
}
