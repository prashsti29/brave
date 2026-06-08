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

    // debug: print loaded env values (avoid printing password in real apps)
    log.Printf("DB envs: DB_HOST=%q DB_PORT=%q DB_USER=%q DB_NAME=%q\n",
        os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_NAME"))

    // debug: also print raw .env file contents as seen by the process
    if data, err := os.ReadFile(".env"); err != nil {
        log.Printf("Failed to read .env: %v", err)
    } else {
        log.Printf(".env raw: %q", string(data))
    }
    connStr := fmt.Sprintf(
        "host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
        os.Getenv("DB_HOST"),
        os.Getenv("DB_PORT"),
        os.Getenv("DB_USER"),
        os.Getenv("DB_PASSWORD"),
        os.Getenv("DB_NAME"),
    )

    db, err := sqlx.Open("postgres", connStr)
    if err != nil {
        log.Fatal("Error connecting to database: ", err)
    }

    err = db.Ping()
    if err != nil {
        log.Fatal("Cannot reach database: ", err)
    }

    log.Println("Database connected successfully")
    return db
}