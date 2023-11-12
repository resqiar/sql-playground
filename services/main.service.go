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
