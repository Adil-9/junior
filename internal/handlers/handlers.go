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
			// logger.DebugLog.Println()
			return
		}
		id = temp
	}
	if id != 0 {
		data := getById(h.DB, id)
		jsonData, err := json.MarshalIndent(data, "", "\t")
		if err != nil {
			logger.DebugLog.Println(err)
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
			// logger.DebugLog.Println()
			return
		}
		pagination = temp
	}

	name := cases.Title(language.Und).String(values.Get("name"))

	ageStr := values.Get("agef")
	var ageF int
	if ageStr != "" {
		temp, err := strconv.Atoi(ageStr)
		if err != nil {
			http.Error(w, "Invalid agef value", http.StatusBadRequest)
			// logger.DebugLog.Println()
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
			// logger.DebugLog.Println()
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
		logger.DebugLog.Println("Error marshaling:", err)
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
		// logger.DebugLog.Println()
		return
	}

	var personFullData structures.PersonFullData
	personFullData.Person = person

	personFullData = requests.GetPersonInfoAPI(person.Name, person.Surname, person.Patronymic)

	// var wg sync.WaitGroup
	// if personFullData.Gender == "" {
	// 	wg.Add(1)
	// 	go requests.GetGender(person.Name, &personFullData, &wg)
	// }
	// if personFullData.Age == 0 {
	// 	wg.Add(1)
	// 	go requests.GetAge(person.Name, &personFullData, &wg)
	// }
	// if personFullData.Country == "" {
	// 	wg.Add(1)
	// 	go requests.GetAge(person.Name, &personFullData, &wg)
	// }

	// wg.Wait()

	data, err := json.MarshalIndent(person, "", "\t")
	if err != nil {
		logger.DebugLog.Println("Error marshaling:", err) // debug?
		// http.Error(w, "Internal server error", http.StatusInternalServerError)
	} else {
		w.Write(data)
	}

	personExists := getIfExists(h.DB, personFullData.Person.Name, personFullData.Person.Surname, personFullData.Person.Patronymic, personFullData.Gender, personFullData.Country, personFullData.Age)
	if personExists {
		w.Write([]byte("Person instance exist"))
		return
	}

	err = addPersonToDB(h.DB, personFullData)
	if err != nil {
		http.Error(w, "Could not add to database", http.StatusInternalServerError)
		return
	}
	// w.Write(data)
}

func (h Handler) handleDelete(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()

	idString := values.Get("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		http.Error(w, "Incorrect id", http.StatusBadRequest)
		return
	}
	person, err := deleteById(h.DB, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	data, err := json.MarshalIndent(person, "", "\t")
	if err != nil {
		logger.DebugLog.Println("Error marshaling:", err) // debug?
		// http.Error(w, "Internal server error", http.StatusInternalServerError)
	} else {
		w.Write(data)
	}

	w.Write([]byte(fmt.Sprintf("Successfully deleted, id: %d", id)))
}

func (h Handler) handlePatch(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	var anyChange bool
	var id int
	idString := values.Get("id")
	if idString != "" {
		temp, err := strconv.Atoi(idString)
		if err != nil {
			http.Error(w, "Invalid id value", http.StatusBadRequest)
			// logger.DebugLog.Println()
			return
		}
		id = temp
	}
	var person structures.PersonFullData
	if id != 0 {
		temp := getById(h.DB, id)
		if temp.Id != 0 {
			person = temp
		} else {
			http.Error(w, "Incorrect id", http.StatusBadRequest)
			return
		}
	}

	name := cases.Title(language.Und).String(values.Get("name"))
	if name == "" {
		name = person.Person.Name
	} else {
		person.Person.Name = name
		anyChange = true
	}

	surname := cases.Title(language.Und).String(values.Get("surname"))
	if surname == "" {
		surname = person.Person.Surname
	} else {
		person.Person.Surname = surname
		anyChange = true
	}

	patronymic := cases.Title(language.Und).String(values.Get("patronymic"))
	if patronymic == "" {
		patronymic = person.Person.Patronymic
	} else {
		person.Person.Patronymic = patronymic
		anyChange = true
	}

	ageStr := values.Get("age")
	var age int
	if ageStr != "" {
		temp, err := strconv.Atoi(ageStr)
		if err != nil {
			http.Error(w, "Invalid age value", http.StatusBadRequest)
			// logger.DebugLog.Println()
			return
		}
		if temp == 0 {
			age = person.Age
		} else {
			age = temp
			person.Age = age
			anyChange = true
		}
	} else {
		age = person.Age
	}

	gender := strings.ToLower(values.Get("gender"))
	if gender != "male" && gender != "female" || gender == "" {
		gender = person.Gender
	} else {
		person.Gender = gender
		anyChange = true
	}

	country := strings.ToUpper(values.Get("country"))
	if country == "" {
		country = person.Country
	} else {
		person.Country = country
		anyChange = true
	}

	if !anyChange {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	if exists := getIfExists(h.DB, name, surname, patronymic, gender, country, age); exists {
		w.Write([]byte("Instance Exist"))
		return
	}
	err := changePersonData(h.DB, id, name, surname, patronymic, gender, age, country)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data, err := json.MarshalIndent(person, "", "\t")
	if err != nil {
		logger.DebugLog.Println("Error marshaling:", err)
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return
	}
	w.Write(data)
	w.Write([]byte("Data changed successfully"))
}
