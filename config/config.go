package config

import (
	"fmt"
	"log"
	"os"

    "github.com/joho/godotenv"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)

func ConnectDB() *gorm.DB {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	
	maskedPassword := "<redacted>"
	maskedJWT := "<redacted>"
	log.Printf("DB envs: DB_HOST=%q DB_PORT=%q DB_USER=%q DB_NAME=%q DB_PASSWORD=%q JWT_SECRET=%q\n",
		os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_NAME"), maskedPassword, maskedJWT)

	databaseURL := os.Getenv("DATABASE_URL")
	var db *gorm.DB
	if databaseURL != "" {
		db, err = gorm.Open(postgres.Open(databaseURL), &gorm.Config{})
		if err != nil {
			log.Fatal("Error connecting to database via DATABASE_URL: ", err)
		}
	} else {
		connStr := fmt.Sprintf(
			"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			os.Getenv("DB_HOST"),
			os.Getenv("DB_PORT"),
			os.Getenv("DB_USER"),
			os.Getenv("DB_PASSWORD"),
			os.Getenv("DB_NAME"),
		)

		db, err = gorm.Open(postgres.Open(connStr), &gorm.Config{})
		if err != nil {
			log.Fatal("Error connecting to database: ", err)
		}
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("Cannot reach database: ", err)
	}

	log.Println("Database connected successfully")
	return db
}
