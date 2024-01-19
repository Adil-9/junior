package requests

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"junior/api/structures"
	"junior/internal"
	"net/http"
	"sync"
	"time"
)

const (
	Genderize   = "https://api.genderize.io/?name="
	Agify       = "https://api.agify.io/?name="
	Nationalize = "https://api.nationalize.io/?name="
)

func GetPersonInfoAPI(name string, surname string, patronymic string) structures.PersonFullData {
	var personFullData structures.PersonFullData
	// link := "Here should be a link to the api"

	// ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	// defer cancel()

	// req, err := http.NewRequestWithContext(ctx, http.MethodGet, link, nil)
	// if err != nil {
	// 	internal.ErrorLog.Println("Creating for person data error: ", err)
	// 	return personFullData
	// }

	// resp, err := http.DefaultClient.Do(req)
	// if err != nil {
	// 	internal.ErrorLog.Println("Request for person data error: ", err)
	// 	return personFullData
	// }

	// body, err := io.ReadAll(resp.Body)
	// if err != nil {
	// 	internal.ErrorLog.Println("Error reading response.Body: ", err)
	// 	return personFullData
	// }

	// var person structures.Person
	// if err := json.Unmarshal(body, &person); err != nil {
	// 	internal.ErrorLog.Println("Error unmarshalling body: ")
	// 	return personFullData
	// }

	personFullData.Person.Name = name
	personFullData.Person.Surname = surname
	personFullData.Person.Patronymic = patronymic

	var wg sync.WaitGroup
	wg.Add(3)

	go GetGender(name, &personFullData, &wg)
	go GetNation(name, &personFullData, &wg)
	go GetAge(name, &personFullData, &wg)

	wg.Wait()

	return personFullData
}

func GetGender(name string, personFullData *structures.PersonFullData, wg *sync.WaitGroup) {
	defer wg.Done()

	var gender structures.PersonGenderize

	link := fmt.Sprintf("%s%s", Genderize, name)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, link, nil)
	if err != nil {
		internal.ErrorLog.Println("Creating for person data error: ", err)
		return
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		internal.ErrorLog.Println("Request for person data error: ", err)
		return
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		internal.ErrorLog.Println("Error reading response.Body: ", err)
		return
	}

	if err := json.Unmarshal(body, &gender); err != nil {
		internal.ErrorLog.Println("Error unmarshalling body: ")
		return
	}

	if gender.Gender == "male" && gender.Probability < 0.5 {
		gender.Gender = "female"
	} else if gender.Gender == "female" && gender.Probability < 0.5 {
		gender.Gender = "male"
	}
	personFullData.Gender = gender.Gender
}

func GetNation(name string, personFullData *structures.PersonFullData, wg *sync.WaitGroup) {
	defer wg.Done()
	var nation structures.PersonNationalize
	link := fmt.Sprintf("%s%s", Nationalize, name)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, link, nil)
	if err != nil {
		internal.ErrorLog.Println("Creating for person data error: ", err)
		return
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		internal.ErrorLog.Println("Request for person data error: ", err)
		return
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		internal.ErrorLog.Println("Error reading response.Body: ", err)
		return
	}

	if err := json.Unmarshal(body, &nation); err != nil {
		internal.ErrorLog.Println("Error unmarshalling body: ")
		return
	}

	var nationality string
	var MaxProbability float32

	for _, v := range nation.CountryList {
		if MaxProbability < v.Probability {
			MaxProbability = v.Probability
			nationality = v.Country_ID
		}
	}

	personFullData.Nationality = nationality
}

func GetAge(name string, personFullData *structures.PersonFullData, wg *sync.WaitGroup) {
	defer wg.Done()
	var age structures.PersonAgify
	link := fmt.Sprintf("%s%s", Agify, name)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, link, nil)
	if err != nil {
		internal.ErrorLog.Println("Creating for person data error: ", err)
		return
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		internal.ErrorLog.Println("Request for person data error: ", err)
		return
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		internal.ErrorLog.Println("Error reading response.Body: ", err)
		return
	}

	if err := json.Unmarshal(body, &age); err != nil {
		internal.ErrorLog.Println("Error unmarshalling body: ")
		return
	}

	personFullData.Age = age.Age
}
