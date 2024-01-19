package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"junior/api/requests"
	"junior/api/structures"
	"junior/internal/logger"
	"net/http"
	"strconv"
	"strings"
	"sync"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
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
			http.Error(w, "Invalid id value", http.StatusBadRequest)
			// logger.ErrorLog.Println()
			return
		}
		id = temp
	}
	if id != 0 {
		data := getById(h.DB, id)
		jsonData, err := json.MarshalIndent(data, "", "\t")
		if err != nil {
			logger.ErrorLog.Println(err)
			http.Error(w, "Error reading request body", http.StatusInternalServerError)
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
			http.Error(w, "Invalid pagination value", http.StatusBadRequest)
			// logger.ErrorLog.Println()
			return
		}
		pagination = temp
	}

	name := cases.Title(language.Und).String(values.Get("name"))

	ageFromString := values.Get("agef")
	var ageF int
	if ageFromString != "" {
		temp, err := strconv.Atoi(ageFromString)
		if err != nil {
			http.Error(w, "Invalid agef value", http.StatusBadRequest)
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
			http.Error(w, "Invalid aget value", http.StatusBadRequest)
			// logger.ErrorLog.Println()
			return
		}
		ageT = temp
	}

	gender := strings.ToLower(values.Get("gender"))
	if gender != "male" && gender != "female" {
		gender = ""
	}

	country := strings.ToUpper(values.Get("country"))

	person := getPerson(h.DB, limit, pagination, ageF, ageT, name, gender, country)

	data, err := json.MarshalIndent(person, "", "\t")
	if err != nil {
		logger.ErrorLog.Println("Error marshaling:", err)
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return
	}

	w.Write(data)
}

func (h Handler) handlePost(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return
	}

	var person structures.Person

	if err = json.Unmarshal(body, &person); err != nil {
		fmt.Println(body) // need to be removed 					<----------------
		http.Error(w, "Bad request", http.StatusBadRequest)
		// logger.ErrorLog.Println()
		return
	}

	var personFullData structures.PersonFullData
	personFullData.Person = person

	personFullData = requests.GetPersonInfoAPI(person.Name, person.Surname, person.Patronymic)

	var wg sync.WaitGroup
	if personFullData.Gender == "" {
		wg.Add(1)
		go requests.GetGender(person.Name, &personFullData, &wg)
	}
	if personFullData.Age == 0 {
		wg.Add(1)
		go requests.GetAge(person.Name, &personFullData, &wg)
	}
	if personFullData.Country == "" {
		wg.Add(1)
		go requests.GetAge(person.Name, &personFullData, &wg)
	}

	wg.Wait()

	err = addPersonToDB(h.DB, personFullData)
	if err != nil {
		http.Error(w, "Could not add to database", http.StatusInternalServerError)
		return
	}

	data, err := json.MarshalIndent(person, "", "\t")
	if err != nil {
		logger.ErrorLog.Println("Error marshaling:", err)
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return
	}

	w.Write(data)
}

func (h Handler) handleDelete(w http.ResponseWriter, r *http.Request) {

}

func (h Handler) handlePatch(w http.ResponseWriter, r *http.Request) {

}
