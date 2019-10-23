package main

import (
	"net/http"

	"./routes"
	"./util"

	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/kataras/golog"
)

var config util.Config
var databaseConnection util.DatabaseConnection

func main() {
	initMigration()
	initRouter()

	defer databaseConnection.Close()
}

func initMigration() {
	err := config.Read("./config.yaml", &config)
	if err != nil {
		panic(err)
	}
	golog.Info("Connecting to database...")
	databaseConnection.Open(config)
	golog.Info("Connected!")
	// Database Drop
	// golog.Warn("Dropping database...")
	// databaseConnection.DropDb()
	// golog.Info("Dropped!")

	// Database Migration
	golog.Warn("Migrating database...")
	databaseConnection.MigrateDb()
	golog.Info("Migrated!")
}

func initRouter() {
	routes := routes.GetAllRoutes(&databaseConnection, config)
	http.Handle("/", routes)
	golog.Info("Server listening at http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		golog.Fatal(err)
		panic(err)
	}
}
