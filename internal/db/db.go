package db

import (
	"database/sql"
	"fmt"
	"junior/internal/logger"

	_ "github.com/lib/pq"
)

// const ( //changeble
// 	host     = "localhost"
// 	port     = 3000
// 	user     = "postgres"
// 	password = "password"
// 	dbname   = "junior"
// )

func InitDB() *sql.DB {

	host := "localhost"    //os.Getenv("DB_HOST")
	port := 3000           //os.Getenv("DB_PORT")
	user := "postgres"     //os.Getenv("DB_USER")
	password := "password" //os.Getenv("DB_PASSWORD")
	dbname := "junior"     //os.Getenv("DB_NAME")
	// sslmode := "disable" os.Getenc("DB_SSL_MODE")

	psqlconn := fmt.Sprintf("host= %s port= %d user= %s password=%s dbname= %s sslmode=disable", host, port, user, password, dbname)
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

func createTables(db *sql.DB) error {
	query := `
		CREATE TABLE IF NOT EXISTS person (
			id SERIAL PRIMARY KEY,
			name TEXT NOT NULL,
			surname TEXT NOT NULL,
			patronymic TEXT NOT NULL,
			age INTEGER NOT NULL,
			gender TEXT NOT NULL,
			country TEXT NOT NULL
		);	
	`
	if _, err := db.Exec(query); err != nil {
		return err
	}
	logger.InfoLog.Println("Created tables in database")
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
