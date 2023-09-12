package main

import (
	"github.com/projects/person-crud/initializers"
	"gorm.io/gorm"
	"log"
	"os"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.InitializeDatabaseConnection()
}

func main() {
	err := CleanUpDatabase()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Migration run complete")
}

func CleanUpDatabase() error {
	database := os.Getenv("DB_DATABASE")
	if err := initializers.DB.Transaction(func(tx *gorm.DB) error {
		// Disable foreign key checks (if supported by your database)
		tx.Exec("SET FOREIGN_KEY_CHECKS = 0")

		// Get all table names
		var tableNames []string
		tx.Raw("SHOW TABLES").Pluck("Tables_in_"+database, &tableNames)

		log.Println(tableNames)
		// Truncate each table
		//for _, tableName := range tableNames {
		//	tx.Exec("TRUNCATE TABLE " + tableName)
		//}

		// Re-enable foreign key checks (if supported by your database)
		tx.Exec("SET FOREIGN_KEY_CHECKS = 1")

		return nil
	}); err != nil {
		panic("Connection Failed")
		return err
	}

	return nil
}
