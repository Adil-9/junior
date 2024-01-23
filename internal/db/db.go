package db

import (
	"database/sql"
	"fmt"
	"junior/internal/logger"
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
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

	psqlconn := fmt.Sprintf("host= %s port= %s user= %s password=%s dbname= %s sslmode=%s", host, port, user, password, dbname, sslmode)
	db, err := sql.Open("postgres", psqlconn)
	if err != nil {
		logger.DebugLog.Fatalln(err)
	}

	err = createTables(db)
	if err != nil {
		logger.DebugLog.Fatalln(err)
	}
	logger.InfoLog.Println("Connected to database")
	return db
}

// func createTables(db *sql.DB) error {
// 	query := `
// 		CREATE TABLE IF NOT EXISTS person (
// 			id SERIAL PRIMARY KEY,
// 			name TEXT NOT NULL,
// 			surname TEXT NOT NULL,
// 			patronymic TEXT NOT NULL,
// 			age INTEGER NOT NULL,
// 			gender TEXT NOT NULL,
// 			country TEXT NOT NULL
// 		);
// 	`
// 	if _, err := db.Exec(query); err != nil {
// 		return err
// 	}
// 	logger.InfoLog.Println("Created tables in database")
// 	return nil
// }

func createTables(db *sql.DB) error {
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		logger.ErrorLog.Fatal(err)
		return err
	}

	m, err := migrate.NewWithDatabaseInstance(
		"migration", // Path to your migration files
		"postgres",  // Database driver name
		driver,
	)
	if err != nil {
		fmt.Println(err)
		logger.ErrorLog.Fatal(err)
		return err
	}

	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		logger.ErrorLog.Fatal(err)
		return err
	}

	fmt.Println("Migrations applied successfully!")
	return nil
}

// func OpenSqliteDB(dbName string) *sql.DB {
// 	db, err := sql.Open("sqlite3", fmt.Sprintf("./%s", dbName))
// 	if err != nil {
// 		logger.DebugLog.Fatal(err)
// 		return nil
// 	}

// 	if err = db.Ping(); err != nil {
// 		logger.DebugLog.Fatal(err)
// 		return nil
// 	}

// 	if err = createTables(db); err != nil {
// 		logger.DebugLog.Fatal(err)
// 		return nil
// 	}

// 	return db
// }
