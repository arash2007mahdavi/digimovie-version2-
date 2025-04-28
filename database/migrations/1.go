package migrations

import (
	"digimovie/src/database"
	"digimovie/src/database/models"
	"gorm.io/gorm"
)

func AddTables() {
	database := database.GetDB()
	
	director := models.Director{}
	movie := models.Movie{}
	user := models.User{}
	
	tables := []interface{}{}
	checkExistsTable(database, director, &tables)
	checkExistsTable(database, movie, &tables)
	checkExistsTable(database, user, &tables)

	err := database.Migrator().CreateTable(tables...)
	if err != nil {
		panic(err)
	}
}

func checkExistsTable(database *gorm.DB, item interface{}, tables *[]interface{}) {
	if !database.Migrator().HasTable(item) {
		*tables = append(*tables, item)
	}
}