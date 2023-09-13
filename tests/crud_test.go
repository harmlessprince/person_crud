package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/projects/person-crud/controllers"
	"github.com/projects/person-crud/initializers"
	"github.com/projects/person-crud/models"
	"github.com/projects/person-crud/utils"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"os"
	"strconv"
	"testing"
)

var Router *gin.Engine

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
	initializers.MigrateFresh()
	fmt.Printf("\033[1;33m%s\033[0m", "> Teardown completed")
	fmt.Printf("\n")
}

func setUpDatabase() {
	os.Setenv("APP_ENV", "test")
	os.Setenv("DB_CONNECTION", "sqlite")
	initializers.InitializeDatabaseConnection()
}

// Mock the database and initialize the Gin router
func setupTestRouter() {
	router := gin.Default()
	router.POST("/api", controllers.StorePerson)
	router.GET("/api", controllers.IndexPerson)
	router.GET("/api/:user_id", controllers.ShowPerson)
	router.PATCH("/api/:user_id", controllers.UpdatePerson)
	router.DELETE("/api/:user_id", controllers.DeletePerson)
	Router = router
}
func TestListingOfPerson(t *testing.T) {

	t.Run("Created Users Are Listed", func(t *testing.T) {

		createPerson("Test Person 1")
		createPerson("Test Person 2")
		createPerson("Test Person 3")
		req, _ := http.NewRequest("GET", "/api", nil)
		req.Header.Set("Accept", "application/json")
		resp := httptest.NewRecorder()

		// Perform the request
		Router.ServeHTTP(resp, req)

		// Check the response status code
		assert.Equal(t, http.StatusOK, resp.Code)

		//Parse the response body into JSONSuccessResult
		var response JSONSuccessResult
		if err := response.ParseSuccessResponseBody(resp.Body.Bytes()); err != nil {
			t.Fatalf("Error parsing response body: %v", err)
		}

		// Replace with appropriate assertions based on your expected results
		assert.Equal(t, "All persons retrieved", response.Message)
		assert.NotEmpty(t, response.Data) // Check if response data is not null
		assert.Equal(t, true, response.Success)

		// Parse the response body into a map
		var responseInMap map[string]interface{}
		if err := json.Unmarshal(resp.Body.Bytes(), &responseInMap); err != nil {
			t.Fatalf("Error parsing response body: %v", err)
		}
		// Checking the "data" field in the response
		data, ok := responseInMap["data"].([]interface{})
		if !ok {
			t.Fatal("Data field should be a list")
		}

		assert.Len(t, data, 3)
	})
}
func TestStorePerson(t *testing.T) {

	t.Run("Valid request", func(t *testing.T) {
		// Create a request with valid request body
		requestBody := []byte(`{"name": "John Doe"}`)
		req, _ := http.NewRequest("POST", "/api", bytes.NewBuffer(requestBody))
		req.Header.Set("Content-Type", "application/json")
		resp := httptest.NewRecorder()

		// Perform the request
		Router.ServeHTTP(resp, req)

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

		assert.NotEmpty(t, person["ID"])
		assert.Equal(t, "John Doe", person["Name"])
	})

	t.Run("Invalid request", func(t *testing.T) {
		//an  invalid request body (e.g., missing "name" field)
		requestBody := []byte(`{"invalid_field": "John Doe"}`)
		req, _ := http.NewRequest("POST", "/api", bytes.NewBuffer(requestBody))
		req.Header.Set("Content-Type", "application/json")
		resp := httptest.NewRecorder()

		// make the request
		Router.ServeHTTP(resp, req)

		// Check the response status code
		assert.Equal(t, http.StatusUnprocessableEntity, resp.Code)

		//Parse the response body into JSONFailedValidationResult
		var response JSONFailedValidationResult
		if err := response.ParseValidationResponseBody(resp.Body.Bytes()); err != nil {
			t.Fatalf("Error parsing response body: %v", err)
		}

		assert.Equal(t, "Validation failed", response.Message)
		assert.Equal(t, false, response.Success)

		expectedErrors := []utils.ValidationError{
			{Condition: "required", Error: "name field is required", Key: "name"},
			// We can add more errors
		}

		assert.ElementsMatch(t, expectedErrors, response.Errors)
	})
}

func TestUpdatePerson(t *testing.T) {

	t.Run("Valid request", func(t *testing.T) {

		person, _ := createPerson("Test Person")
		// Create a request with valid request body
		requestBody := []byte(`{"name": "John Doe"}`)
		req, _ := http.NewRequest("PATCH", "/api/"+strconv.Itoa(int(person.ID)), bytes.NewBuffer(requestBody))
		req.Header.Set("Content-Type", "application/json")
		resp := httptest.NewRecorder()

		// Perform the request
		Router.ServeHTTP(resp, req)

		// Check the response status code
		assert.Equal(t, http.StatusOK, resp.Code)

		//Parse the response body into JSONSuccessResult
		var response JSONSuccessResult
		if err := response.ParseSuccessResponseBody(resp.Body.Bytes()); err != nil {
			t.Fatalf("Error parsing response body: %v", err)
		}

		// Replace with appropriate assertions based on your expected results
		assert.Equal(t, "Person updated successfully", response.Message)
		assert.NotEmpty(t, response.Data) // Check if the person was created with an ID
		assert.Equal(t, true, response.Success)

		personFromResponse, _ := response.Data.(map[string]interface{})
		personFromResponseId := personFromResponse["ID"]
		floatValue, _ := personFromResponseId.(float64)
		assert.Equal(t, strconv.Itoa(int(person.ID)), strconv.FormatFloat(floatValue, 'f', -1, 64))
		assert.Equal(t, "John Doe", personFromResponse["Name"].(string))
	})

	t.Run("Invalid request", func(t *testing.T) {
		//an  invalid request body (e.g., missing "name" field)
		requestBody := []byte(`{"invalid_field": "John Doe"}`)
		req, _ := http.NewRequest("PATCH", "/api/100", bytes.NewBuffer(requestBody))
		req.Header.Set("Content-Type", "application/json")
		resp := httptest.NewRecorder()

		// make the request
		Router.ServeHTTP(resp, req)

		// Check the response status code
		assert.Equal(t, http.StatusNotFound, resp.Code)

		//Parse the response body into JSONErrorResult
		var response JSONErrorResult
		if err := response.ParseErrorResponseBody(resp.Body.Bytes()); err != nil {
			t.Fatalf("Error parsing response body: %v", err)
		}

		assert.Equal(t, "Person could not be found", response.Message)
		assert.Equal(t, false, response.Success)

	})
}

func TestShowPerson(t *testing.T) {

	t.Run("Valid User ID request", func(t *testing.T) {

		person, _ := createPerson("Test Person")
		req, _ := http.NewRequest("GET", "/api/"+strconv.Itoa(int(person.ID)), nil)
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Accept", "application/json")
		resp := httptest.NewRecorder()

		// Perform the request
		Router.ServeHTTP(resp, req)

		// Check the response status code
		assert.Equal(t, http.StatusOK, resp.Code)

		//Parse the response body into JSONSuccessResult
		var response JSONSuccessResult
		if err := response.ParseSuccessResponseBody(resp.Body.Bytes()); err != nil {
			t.Fatalf("Error parsing response body: %v", err)
		}

		// Replace with appropriate assertions based on your expected results
		assert.Equal(t, "Person retrieved", response.Message)
		assert.NotEmpty(t, response.Data) // Check if response data is not null
		assert.Equal(t, true, response.Success)

		personFromResponse, _ := response.Data.(map[string]interface{})
		personFromResponseId := personFromResponse["ID"]
		floatValue, _ := personFromResponseId.(float64)
		assert.Equal(t, strconv.Itoa(int(person.ID)), strconv.FormatFloat(floatValue, 'f', -1, 64))
		assert.Equal(t, "Test Person", personFromResponse["Name"].(string))
	})

	t.Run("Invalid ID Supplied", func(t *testing.T) {
		//an  invalid request body (e.g., missing "name" field)
		createPerson("Test Person")
		req, _ := http.NewRequest("GET", "/api/1000", nil)
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Accept", "application/json")
		resp := httptest.NewRecorder()

		// Perform the request
		Router.ServeHTTP(resp, req)

		// Check the response status code
		assert.Equal(t, http.StatusNotFound, resp.Code)

		//Parse the response body into JSONErrorResult
		var response JSONErrorResult
		if err := response.ParseErrorResponseBody(resp.Body.Bytes()); err != nil {
			t.Fatalf("Error parsing response body: %v", err)
		}

		assert.Equal(t, "Person could not be found", response.Message)
		assert.Equal(t, false, response.Success)

	})
}

type JSONSuccessResult utils.JSONSuccessResult
type JSONErrorResult utils.JSONErrorResult
type JSONFailedValidationResult utils.JSONFailedValidationResult

func (r *JSONSuccessResult) ParseSuccessResponseBody(body []byte) error {
	if err := json.Unmarshal(body, r); err != nil {
		return err
	}
	return nil
}

func (r *JSONErrorResult) ParseErrorResponseBody(body []byte) error {

	if err := json.Unmarshal(body, r); err != nil {
		return err
	}
	return nil
}

func (r *JSONFailedValidationResult) ParseValidationResponseBody(body []byte) error {
	if err := json.Unmarshal(body, r); err != nil {
		return err
	}
	return nil
}

func createPerson(name string) (models.Person, bool) {
	var person models.Person
	if err := initializers.DB.Where(models.Person{Name: name}).FirstOrCreate(&person).Error; err != nil {
		panic(err)
	}
	return person, true
}
