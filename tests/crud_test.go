package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/projects/person-crud/controllers"
	"github.com/projects/person-crud/initializers"
	"github.com/projects/person-crud/utils"
	"github.com/stretchr/testify/assert"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	teardown()
	os.Exit(code)
}
func setup() {
	setUpDatabase()
	initializers.MigrateUp()
	setupTestRouter()
	fmt.Printf("\033[1;33m%s\033[0m", "> Setup completed\n")
}

func teardown() {
	//setUpDatabase()
	//initializers.MigrateFresh()
	fmt.Printf("\033[1;33m%s\033[0m", "> Teardown completed")
	fmt.Printf("\n")
}

func setUpDatabase() {
	os.Setenv("APP_ENV", "test")
	os.Setenv("DB_CONNECTION", "sqlite")
	initializers.InitializeDatabaseConnection()
}

// Mock the database and initialize the Gin router
func setupTestRouter() *gin.Engine {
	router := gin.Default()
	router.POST("/api", controllers.StorePerson)
	router.GET("/api", controllers.IndexPerson)
	router.GET("/api/:user_id", controllers.ShowPerson)
	router.PATCH("/api/:user_id", controllers.UpdatePerson)
	router.DELETE("/api/:user_id", controllers.DeletePerson)
	return router
}

func TestStorePerson(t *testing.T) {
	// Create a test router
	router := setupTestRouter()

	t.Run("Valid request", func(t *testing.T) {
		// Create a request with valid request body
		requestBody := []byte(`{"name": "John Doe"}`)
		req, _ := http.NewRequest("POST", "/api", bytes.NewBuffer(requestBody))
		req.Header.Set("Content-Type", "application/json")
		resp := httptest.NewRecorder()

		// Perform the request
		router.ServeHTTP(resp, req)

		// Check the response status code
		assert.Equal(t, http.StatusOK, resp.Code)

		//Parse the response body into JSONSuccessResult
		var response JSONSuccessResult
		if err := response.ParseSuccessResponseBody(resp.Body.Bytes()); err != nil {
			t.Fatalf("Error parsing response body: %v", err)
		}

		// Replace with appropriate assertions based on your expected results
		assert.Equal(t, "Person created successfully", response.Message)
		assert.NotEmpty(t, response.Data) // Check if the person was created with an ID
		assert.Equal(t, true, response.Success)

		person, _ := response.Data.(map[string]interface{})
		log.Println(person["Name"])
		assert.NotEmpty(t, person["ID"])
		assert.Equal(t, "John Doe", person["Name"])
	})

	//t.Run("Invalid request", func(t *testing.T) {
	//	// Create a request with an invalid request body (e.g., missing "name" field)
	//	requestBody := []byte(`{"invalid_field": "John Doe"}`)
	//	req, _ := http.NewRequest("POST", "/persons", bytes.NewBuffer(requestBody))
	//	req.Header.Set("Content-Type", "application/json")
	//	resp := httptest.NewRecorder()
	//
	//	// Perform the request
	//	router.ServeHTTP(resp, req)
	//
	//	// Check the response status code
	//	assert.Equal(t, http.StatusBadRequest, resp.Code)
	//
	//	// Decode the response body and assert its content
	//	var validationErrors map[string][]string
	//
	//	// Replace with appropriate assertions based on your expected validation error response
	//	assert.Contains(t, validationErrors, "name")
	//})
}

type JSONSuccessResult utils.JSONSuccessResult

func (r *JSONSuccessResult) ParseSuccessResponseBody(body []byte) error {
	if err := json.Unmarshal(body, r); err != nil {
		return err
	}
	return nil
}
