package main

import (
	"github.com/gin-gonic/gin"
	"github.com/projects/person-crud/controllers"
	"github.com/projects/person-crud/docs"
	"github.com/projects/person-crud/initializers"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
	"os"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.InitializeDatabaseConnection()
	gin.ForceConsoleColor()
}

// @title Swagger Person Crud API
// @version 1.0
// @description "This is a crud server for creating resource."

// @contact.name API Support
// @contact.email realolamilekan@gmail.io

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @host localhost:3000
// @BasePath /api
// @query.collection.format multi
func main() {
	appUrl := os.Getenv("APP_URL")
	appPort := os.Getenv("APP_PORT")
	appEnv := os.Getenv("APP_ENV")

	router := gin.Default()
	// programmatically set swagger info
	docs.SwaggerInfo.Title = "HNGx Stage 2 Person Crud API"
	docs.SwaggerInfo.Description = "This is a crud server for modifying person resource."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = appUrl
	docs.SwaggerInfo.BasePath = "/api"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	router.POST("/", controllers.HelloWorld)
	router.POST("/api", controllers.StorePerson)
	router.GET("/api", controllers.IndexPerson)
	router.GET("/api/:user_id", controllers.ShowPerson)
	router.PATCH("/api/:user_id", controllers.UpdatePerson)
	router.DELETE("/api/:user_id", controllers.DeletePerson)
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	if appEnv == "local" {
		err := router.Run(appUrl) // listen and serve on localhost:3000
		if err != nil {
			log.Panicf("error: %s", err)
		}
	} else {
		err := router.Run(":" + appPort)
		if err != nil {
			log.Panicf("error: %s", err)
		}
	}
}
