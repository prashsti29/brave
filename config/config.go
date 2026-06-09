package config

import (
	"fmt"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func ConnectDB() *sqlx.DB {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Print non-sensitive loaded env values; do not print raw .env or secrets.
	maskedPassword := "<redacted>"
	maskedJWT := "<redacted>"
	log.Printf("DB envs: DB_HOST=%q DB_PORT=%q DB_USER=%q DB_NAME=%q DB_PASSWORD=%q JWT_SECRET=%q\n",
		os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_NAME"), maskedPassword, maskedJWT)

	// Allow overriding connection via DATABASE_URL (e.g. postgres://user:pass@host:port/db)
	databaseURL := os.Getenv("DATABASE_URL")
	var db *sqlx.DB
	if databaseURL != "" {
		db, err = sqlx.Open("postgres", databaseURL)
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

		db, err = sqlx.Open("postgres", connStr)
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
