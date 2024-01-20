package handlers

import (
	"context"
	"database/sql"
	"fmt"
	"junior/api/structures"
	"junior/internal/logger"
	"time"
)

func getIfExists(db *sql.DB, name, surname, patronymic, gender, country string, age int) bool {
	var person structures.PersonFullData
	query := fmt.Sprintf(`
	SELECT ID, Name, Surname, Patronymic, Age, Gender, Country
	FROM PERSON
	WHERE name = '%s' AND surname = '%s' AND patronymic = '%s' AND age = (%d) AND gender = '%s' AND country = '%s';
	`, name, surname, patronymic, age, gender, country)
	row := db.QueryRow(query)
	err := row.Scan(&person.Id, &person.Person.Name, &person.Person.Surname, &person.Person.Patronymic, &person.Age, &person.Gender, &person.Country)
	if err != nil {
		logger.ErrorLog.Println("Error retrieving information from data base: ", err)
		return false
	}
	return true
}

func deleteById(db *sql.DB, id int) (structures.PersonFullData, error) {
	var person structures.PersonFullData

	personTemp := getById(db, id)
	if personTemp.Id == 0 {
		return person, fmt.Errorf("no person instance with such id: %d", id)
	}

	query := fmt.Sprintf(`
	DELETE FROM person
	WHERE id = %d
	RETURNING *;
	`, id)

	row, err := db.Query(query)
	if err != nil {
		logger.DebugLog.Printf("Could not delete from database id: %d, %v", id, err)
		return person, fmt.Errorf("internal server error")
	}
	err = row.Scan(&person.Id, &person.Person.Name, &person.Person.Surname, &person.Person.Patronymic, &person.Age, &person.Gender, &person.Country)
	if err != nil {
		logger.DebugLog.Println("Error retrieving information from data base: ", err)
	}
	return person, nil
}

// june 18

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

func addPersonToDB(db *sql.DB, person structures.PersonFullData) error {
	queryInsert := fmt.Sprintf(`
	INSERT INTO person (name, surname, patronymic, age, gender, country) VALUES
  		('%s', '%s', '%s', %d, '%s', '%s');
	`, person.Person.Name, person.Person.Surname, person.Person.Patronymic, person.Age, person.Gender, person.Country)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*7)
	defer cancel()

	fmt.Println(person.Person)
	fmt.Println(person)

	_, err := db.ExecContext(ctx, queryInsert)
	if err != nil {
		logger.DebugLog.Println("Could not insert into database:", err)
		return err
	}
	return nil
}
