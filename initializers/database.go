package initializers

import (
	"fmt"
	"github.com/TwiN/go-color"
	"github.com/projects/person-crud/models"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"os"
)

var DB *gorm.DB

func InitializeDatabaseConnection() {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details

	var err error
	driver := os.Getenv("DB_CONNECTION")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	database := os.Getenv("DB_DATABASE")
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	ssl_mode := os.Getenv("DB_SSL_MODE")
	var dsn string
	if driver == "mysql" {
		dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, host, port, database)
		DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	}

	if driver == "postgres" {
		dsn = fmt.Sprintf(
			"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
			host, username, password, database, port, ssl_mode,
		)
		DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	}
	if driver == "sqlite" {
		DB, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	}
	if err != nil {
		panic("Failed to connect to database")
	}

	log.Println(color.Ize(color.Green, "Database connection successful"))

}

func MigrateUp() {
	// Migrate the schema
	err := DB.AutoMigrate(&models.Person{})
	if err != nil {
		log.Fatal(err)
	}
	log.Println(color.Ize(color.Green, "Migration up complete"))
}

func MigrateFresh() {
	// Migrate the schema
	err := DB.Migrator().DropTable(&models.Person{})
	if err != nil {
		return
	}
	MigrateUp()
	log.Println(color.Ize(color.Green, "Fresh Migration run complete"))
}
