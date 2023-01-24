package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func EnvUsername() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	return os.Getenv("USERNAME")
}

func EnvPassword() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	return os.Getenv("PASSWORD")
}

func EnvDBName() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	return os.Getenv("DBNAME")
}

