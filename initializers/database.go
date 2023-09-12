package initializers

import (
	"fmt"
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
	var dsn string
	if driver == "mysql" {
		dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, host, port, database)
		DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	}

	if driver == "postgres" {
		dsn = fmt.Sprintf(
			"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
			host, username, password, database, port,
		)
		DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	}
	if driver == "sqlite" {
		DB, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	}

	if err != nil {
		panic("Failed to connect to database")
	}
	log.Println("Database connection successful")
}
