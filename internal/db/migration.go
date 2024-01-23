package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func migrate(db *sql.DB) {
	loadEnv()

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	sslmode := os.Getenv("DB_SSL_MODE")

	psqlconn := fmt.Sprintf("host= %s port= %s user= %s password=%s dbname= %s sslmode=%s", host, port, user, password, dbname, sslmode)

	err := runMigrations(db, migrateDir)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Migrations completed successfully.")
}

func runMigrations(db *sql.DB, migrationDir string) error {
	m, err := migrate.New(
		migrationDir,
		"postgres://"+dbSource,
	)
	if err != nil {
		return err
	}

	// Run the migrations
	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		return err
	}

	return nil
}
