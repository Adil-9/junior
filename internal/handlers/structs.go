package handlers

import (
	"database/sql"
	"junior/internal/db"
)

type Handler struct {
	DB *sql.DB
}

func CreateHandler() Handler {
	db := db.InitDB()
	return Handler{
		DB: db,
	}
}
