package main

import (
	"digimovie/src/database"
	"digimovie/src/config"
	"digimovie/src/database/migrations"
	"digimovie/src/server"
)

func main() {
	cfg := config.GetConfig()
	database.InitRedis(cfg)
	defer database.CLoseRedis()
	err := database.InitDB(cfg)
	if err != nil {
		panic(err)
	}
	defer database.CloseDB()
	migrations.AddTables()
	server.InitServer(cfg)
}