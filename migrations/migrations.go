package main

import (
	"github.com/projects/person-crud/initializers"
	"github.com/projects/person-crud/models"
	"log"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.InitializeDatabaseConnection()
}

func main() {
	// Migrate the schema
	err := initializers.DB.AutoMigrate(&models.Person{})
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Migration run complete")
}
