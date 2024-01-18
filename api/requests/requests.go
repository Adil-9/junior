package requests

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"junior/api/structures"
	"junior/internal"
	"net/http"
	"time"
)

const (
	Genderize   = "https://api.genderize.io/?name="
	Agify       = "https://api.agify.io/?name="
	Nationalize = "https://api.nationalize.io/?name="
)

func GetPersonInfoAPI(name string, surname string, patronymic string) {
	link := "Here should be a link to the api"

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, link, nil)
	if err != nil {
		internal.ErrorLog.Println("Creating for person data error: ", err)
		//doiferror
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		internal.ErrorLog.Println("Request for person data error: ", err)
		//doiferror
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		internal.ErrorLog.Println("Error reading response.Body: ", err)
		//doiferror
	}

	var person structures.Person
	if err := json.Unmarshal(body, &person); err != nil {
		internal.ErrorLog.Println("Error unmarshalling body: ")
		//doiferror
	}
	if person.Name == "" {
		//doiferror
		return
	}

	var personFullData structures.PersonFullData

	personFullData.Person.Name = name
	personFullData.Person.Surname = surname
	personFullData.Person.Patronymic = patronymic

	personFullData.Gender = GetGender(name)
	personFullData.Nationality = GetNation(name)
	personFullData.Age = GetAge(name)
}

func GetGender(name string) structures.PersonGenderize {
	var gender structures.PersonGenderize
	link := fmt.Sprintf("%s%s", Genderize, name)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, link, nil)
	if err != nil {
		internal.ErrorLog.Println("Creating for person data error: ", err)
		return gender
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		internal.ErrorLog.Println("Request for person data error: ", err)
		return gender
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		internal.ErrorLog.Println("Error reading response.Body: ", err)
		return gender
	}

	if err := json.Unmarshal(body, &gender); err != nil {
		internal.ErrorLog.Println("Error unmarshalling body: ")
		return gender
	}

	return gender
}

func GetNation(name string) structures.PersonNationalize {
	var nation structures.PersonNationalize
	link := fmt.Sprintf("%s%s", Nationalize, name)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, link, nil)
	if err != nil {
		internal.ErrorLog.Println("Creating for person data error: ", err)
		return nation
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		internal.ErrorLog.Println("Request for person data error: ", err)
		return nation
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		internal.ErrorLog.Println("Error reading response.Body: ", err)
		return nation
	}

	if err := json.Unmarshal(body, &nation); err != nil {
		internal.ErrorLog.Println("Error unmarshalling body: ")
		return nation
	}

	return nation
}
func GetAge(name string) structures.PersonAgify {
	var age structures.PersonAgify
	link := fmt.Sprintf("%s%s", Agify, name)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, link, nil)
	if err != nil {
		internal.ErrorLog.Println("Creating for person data error: ", err)
		return age
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		internal.ErrorLog.Println("Request for person data error: ", err)
		return age
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		internal.ErrorLog.Println("Error reading response.Body: ", err)
		return age
	}

	if err := json.Unmarshal(body, &age); err != nil {
		internal.ErrorLog.Println("Error unmarshalling body: ")
		return age
	}

	return age
}
