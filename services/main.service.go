package services

import (
	"log"
	"sql-playground/db"
)

type World struct {
	ID          int
	Name        string
	CountryCode string
	District    string
	Population  int
}

func GetAll() []World {
	var worlds []World

	SQL := "SELECT * FROM city LIMIT 10"

	rows, err := db.Raw.Query(SQL)
	if err != nil {
		log.Println(err)
		return nil
	}
	defer rows.Close()

	for rows.Next() {
		var world World

		if err := rows.Scan(
			&world.ID,
			&world.Name,
			&world.CountryCode,
			&world.District,
			&world.Population,
		); err != nil {
			log.Println(err)
			return nil
		}

		worlds = append(worlds, world)
	}

	return worlds
}
