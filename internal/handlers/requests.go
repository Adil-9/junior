package handlers

import (
	"database/sql"
	"fmt"
	"junior/api/structures"
	"junior/internal/logger"
)

func getById(db *sql.DB, id int) structures.PersonFullData {
	var person structures.PersonFullData
	query := `
	SELECT ID, Name, Surname, Patronymic, Age, Gender, Country
	FROM PERSON
	WHERE id = ($1);
	`
	row := db.QueryRow(query, id)
	err := row.Scan(&person.Id, &person.Person.Name, &person.Person.Surname, &person.Person.Patronymic, &person.Age, &person.Gender, &person.Country)
	if err != nil {
		logger.ErrorLog.Println("Error retrieving information from data base: ", err)
	}

	return person
}

func getPerson(db *sql.DB, limit, pagination, ageF, ageT int, name, gender, country string) []structures.PersonFullData {
	var people []structures.PersonFullData
	queryDyn := `
	SELECT ID, Name, Surname, Patronymic, Age, Gender, Country
	FROM PERSON WHERE 1=1
	`
	
	if name != "" {
		queryDyn += fmt.Sprintf(` AND Name = '%s' `, name)
	}
	queryDyn += fmt.Sprintf(` AND Age >= %d `, ageF)
	if ageT != 0 {
		queryDyn += fmt.Sprintf(` AND Age <= %d `, ageT)
	}
	if gender != "" {
		queryDyn += fmt.Sprintf(` AND Gender = '%s' `, gender)
	}
	if country != "" {
		queryDyn += fmt.Sprintf(` AND Country = '%s' `, country)
	}
	queryDyn += fmt.Sprintf(` LIMIT %d OFFSET %d`, limit, (pagination-1)*limit)

	rows, err := db.Query(queryDyn)
	if err != nil {
		logger.ErrorLog.Println("Error retrieving data from data base: ", err)
		return people
	}
	defer rows.Close()

	for rows.Next() {
		var person structures.PersonFullData
		err := rows.Scan(&person.Id, &person.Person.Name, &person.Person.Surname, &person.Person.Patronymic, &person.Age, &person.Gender, &person.Country)
		if err != nil {
			logger.ErrorLog.Println("Error retrieving data from data base: ", err)
		}
		people = append(people, person)
	}

	return people
}
