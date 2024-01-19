package handlers

import (
	"encoding/json"
	"junior/internal/logger"
	"net/http"
	"strconv"
)

func (h Handler) HandleRequest(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.handleGet(w, r)
	case http.MethodPost:
		h.handlePost(w, r)
	case http.MethodDelete:
		h.handleDelete(w, r)
	case http.MethodPatch:
		h.handlePatch(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (h Handler) handleGet(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()

	var id int
	idString := values.Get("id")
	if idString != "" {
		temp, err := strconv.Atoi(idString)
		if err != nil {
			w.Write([]byte("Invalid id"))
			// logger.ErrorLog.Println()
			return
		}
		id = temp
	}
	if id != 0 {
		data := getById(h.DB, id)
		jsonData, err := json.Marshal(data)
		if err != nil {
			logger.ErrorLog.Println(err)
			w.Write([]byte("Internal server error"))
			return
		}
		w.Write(jsonData)
		return
	}

	var limit int = 50
	limitString := values.Get("limit")
	if limitString != "" {
		temp, err := strconv.Atoi(limitString)
		if err == nil {
			limit = temp
		}
	}

	var pagination int = 1
	paginationString := values.Get("pagination")
	if paginationString != "" {
		temp, err := strconv.Atoi(paginationString)
		if err != nil {
			w.Write([]byte("Invalid pagination"))
			// logger.ErrorLog.Println()
			return
		}
		pagination = temp
	}

	name := values.Get("name")

	ageFromString := values.Get("agef")
	var ageF int
	if ageFromString != "" {
		temp, err := strconv.Atoi(ageFromString)
		if err != nil {
			w.Write([]byte("Invalid start agef value"))
			// logger.ErrorLog.Println()
			return
		}
		ageF = temp
	}

	ageToString := values.Get("aget")
	var ageT int
	if ageToString != "" {
		temp, err := strconv.Atoi(ageToString)
		if err != nil {
			w.Write([]byte("Invalid aget value"))
			// logger.ErrorLog.Println()
			return
		}
		ageT = temp
	}

	gender := values.Get("gender")
	if gender != "male" && gender != "female" {
		gender = ""
	}

	country := values.Get("country")

	person := getPerson(h.DB, limit, pagination, ageF, ageT, name, gender, country)

	data, err := json.Marshal(person)
	if err != nil {
		logger.ErrorLog.Println("Error marshaling:", err)
		w.Write([]byte("Internal server error"))
	}
	w.Write(data)
}

func (h Handler) handlePost(w http.ResponseWriter, r *http.Request) {

}

func (h Handler) handleDelete(w http.ResponseWriter, r *http.Request) {

}

func (h Handler) handlePatch(w http.ResponseWriter, r *http.Request) {

}
