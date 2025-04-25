package migrations

import (
	"digimovie/src/database"
	"digimovie/src/database/models"
)

func AddTables() {
	database := database.GetDB()
	
	tables := []interface{}{}

	movie := models.Movie{}
	if !database.Migrator().HasTable(movie) {
		tables = append(tables, movie)
	}

	database.Migrator().CreateTable(tables...)
}