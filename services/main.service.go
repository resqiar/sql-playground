package services

import (
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

func GetTotalCountry() (int, error) {
	SQL := "SELECT COUNT(DISTINCT countrycode) FROM city;"

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

func GetTotalDistrict() (int, error) {
	SQL := "SELECT COUNT(DISTINCT district) FROM city;"

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
