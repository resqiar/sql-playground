package services

import (
	"fmt"
	"log"
	"sql-playground/db"
)

type City struct {
	ID          int
	Name        string
	CountryCode string
	District    string
	Population  int
}

type CountryCapital struct {
	Code    string
	Name    string
	Capital string
}

var (
	allowedColumns = [5]string{"id", "name", "countrycode", "district", "population"}
)

func GetAllCities(page int) []City {
	var cities []City

	SQL := "SELECT * FROM city LIMIT ? OFFSET ?;"
	pageSize := 100
	offset := (page - 1) * pageSize

	rows, err := db.Raw.Query(SQL, pageSize, offset)
	if err != nil {
		log.Println(err)
		return nil
	}
	defer rows.Close()

	for rows.Next() {
		var city City

		if err := rows.Scan(
			&city.ID,
			&city.Name,
			&city.CountryCode,
			&city.District,
			&city.Population,
		); err != nil {
			log.Println(err)
			return nil
		}

		cities = append(cities, city)
	}

	return cities
}

func GetCountCities() (int, error) {
	SQL := "SELECT COUNT(ID) FROM city;"

	row := db.Raw.QueryRow(SQL)
	if row.Err() != nil {
		log.Println(row.Err())
		return 0, row.Err()
	}

	var count int
	if err := row.Scan(&count); err != nil {
		log.Println(row.Err())
		return 0, row.Err()
	}

	return count, nil
}

func GetSumPopulation() (int, error) {
	SQL := "SELECT SUM(population) FROM city;"

	row := db.Raw.QueryRow(SQL)
	if row.Err() != nil {
		log.Println(row.Err())
		return 0, row.Err()
	}

	var count int
	if err := row.Scan(&count); err != nil {
		log.Println(row.Err())
		return 0, row.Err()
	}

	return count, nil
}

func GetTotal(input string) (int, error) {
	var validInput = false
	for _, v := range allowedColumns {
		if v == input {
			validInput = true
		}
	}
	if !validInput {
		return 0, nil
	}

	SQL := fmt.Sprintf("SELECT COUNT(DISTINCT %s) FROM city;", input)

	row := db.Raw.QueryRow(SQL)
	if row.Err() != nil {
		log.Println(row.Err())
		return 0, row.Err()
	}

	var count int
	if err := row.Scan(&count); err != nil {
		log.Println(row.Err())
		return 0, row.Err()
	}

	return count, nil
}

func Filter(id, name string, country string, district string) []City {
	SQL := "SELECT * FROM city WHERE id LIKE ? AND name LIKE ? AND countrycode LIKE ? AND district LIKE ? LIMIT 100;"

	var cities []City
	rows, err := db.Raw.Query(SQL, id, name, country, district)
	if err != nil {
		log.Println(err)
		return nil
	}
	defer rows.Close()

	for rows.Next() {
		var city City

		if err := rows.Scan(
			&city.ID,
			&city.Name,
			&city.CountryCode,
			&city.District,
			&city.Population,
		); err != nil {
			log.Println(err)
			return nil
		}

		cities = append(cities, city)
	}

	return cities
}

func GetAllCountryCapital(page int) []CountryCapital {
	var countries []CountryCapital

	SQL := "SELECT country.code, country.name, city.name AS capital FROM country INNER JOIN city ON country.capital = city.id LIMIT ? OFFSET ?;"
	pageSize := 100
	offset := (page - 1) * pageSize

	rows, err := db.Raw.Query(SQL, pageSize, offset)
	if err != nil {
		log.Println(err)
		return nil
	}
	defer rows.Close()

	for rows.Next() {
		var country CountryCapital

		if err := rows.Scan(
			&country.Code,
			&country.Name,
			&country.Capital,
		); err != nil {
			log.Println(err)
			return nil
		}

		countries = append(countries, country)
	}

	return countries
}

func FilterCountry(code string, name string, capital string) []CountryCapital {
	SQL := `
		SELECT country.code, country.name, city.name AS capital
		FROM country INNER JOIN city ON country.capital = city.id
		WHERE country.code LIKE ? AND country.name LIKE ? AND city.name LIKE ?
		LIMIT 100;
	`

	var countries []CountryCapital
	rows, err := db.Raw.Query(SQL, code, name, capital)
	if err != nil {
		log.Println(err)
		return nil
	}
	defer rows.Close()

	for rows.Next() {
		var country CountryCapital

		if err := rows.Scan(
			&country.Code,
			&country.Name,
			&country.Capital,
		); err != nil {
			log.Println(err)
			return nil
		}

		countries = append(countries, country)
	}

	return countries
}
