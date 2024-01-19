package handlers

import "database/sql"

type Handler struct {
	DB *sql.DB
}

func CreateHandler(db *sql.DB) Handler {
	return Handler{
		DB: db,
	}
}