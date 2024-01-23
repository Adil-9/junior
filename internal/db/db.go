package db

import (
	"database/sql"
	"fmt"
	"junior/internal/logger"
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func loadEnv() { // load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
}

func InitDB() *sql.DB {
	loadEnv()

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	sslmode := os.Getenv("DB_SSL_MODE")
	
	url := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s", user, password, host, port, dbname, sslmode)

	err := migrateDatabase(url)
	if err != nil {
		logger.DebugLog.Fatalln(err)
	}

	db, err := sql.Open("postgres", url)
	if err != nil {
		logger.DebugLog.Fatalln(err)
	}

	logger.InfoLog.Println("Connected to database")
	return db
}

func migrateDatabase(url string) error {
	// postgres://postgres:password@postgres:5432/postgres?sslmode=disable
	m, err := migrate.New(
		"file://migration",
		url)
	if err != nil {
		return err
	}
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return err
	}
	return nil
}
