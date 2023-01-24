package config

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DB() *gorm.DB {
	URL := fmt.Sprintf("postgres://%v:%s@localhost:5432/%t", EnvUsername(), EnvPassword(), EnvDBName())

	db, err := gorm.Open(postgres.Open(URL), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	return db
}
