package db

import (
	"database/sql"
	"fmt"
	"junior/internal/logger"

	_ "github.com/lib/pq"
)

const ( //changeble
	host     = "localhost"
	port     = 3000
	user     = "postgres"
	password = "password"
	dbname   = "junior"
)

func InitDB() *sql.DB {
	psqlconn := fmt.Sprintf("host= %s port= %d user= %s password=%s dbname= %s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlconn)
	if err != nil {
		logger.ErrorLog.Fatalln(err)
	}

	err = createTables(db)
	if err != nil {
		logger.ErrorLog.Fatalln(err)
	}

	return db
}

func createTables(db *sql.DB) error {
	query := `
		CREATE TABLE IF NOT EXISTS PERSON (
			ID INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
			Name TEXT NOT NULL,
			Surname TEXT NOT NULL,
			Patronymic TEXT NOT NULL,
			Age INTEGER NOT NULL,
			Gender TEXT NOT NULL,
			Country TEXT NOT NULL
		);
	`
	if _, err := db.Exec(query); err != nil {
		return err
	}
	return nil
}

// func OpenSqliteDB(dbName string) *sql.DB {
// 	db, err := sql.Open("sqlite3", fmt.Sprintf("./%s", dbName))
// 	if err != nil {
// 		logger.ErrorLog.Fatal(err)
// 		return nil
// 	}

// 	if err = db.Ping(); err != nil {
// 		logger.ErrorLog.Fatal(err)
// 		return nil
// 	}

// 	if err = createTables(db); err != nil {
// 		logger.ErrorLog.Fatal(err)
// 		return nil
// 	}

// 	return db
// }
