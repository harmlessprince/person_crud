package controllers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/projects/person-crud/initializers"
	"github.com/projects/person-crud/models"
	"github.com/projects/person-crud/utils"
	"gorm.io/gorm"
	"time"
)

type StorePersonBody struct {
	Name string `json:"name"  binding:"required" example:"Mark Essien"`
}

type UpdatePersonBody struct {
	Name string `json:"name"  binding:"omitempty" example:"Mark Essien"`
}

// PersonDto
// represents data about a person in the people list
type PersonDto struct {
	ID        string    `json:"id" example:"1"`
	Name      string    `json:"name" example:"Mark Essien"`
	CreatedAt time.Time `json:"created_at" example:"2023-09-11T16:23:57.832758+01:00"`
	UpdatedAt time.Time `json:"updated_at" example:"2023-09-11T16:23:57.832758+01:00"`
	DeletedAt time.Time `json:"deleted_at" example:"2023-09-11T16:23:57.832758+01:00"`
}

type FilterPersonQueryParams struct {
	Name string `form:"name"  binding:"omitempty" example:"Mark"`
}

// StorePerson
// @Description store a new person
// @Summary Store User
// @Tags Person CRUD
// @ID store-person
// @Accept json
// @Produce json
// @Param data body StorePersonBody true "Person object to create"
// @Success 200 {object} utils.JSONSuccessResult{data=PersonDto,message}
// @Failure 422 {array} utils.ValidationError
// @Router /person [post]
func StorePerson(context *gin.Context) {
	// Get data of request body
	var body StorePersonBody

	validationErrors := utils.ValidateRequest(context, &body)
	if len(validationErrors) > 0 {
		utils.SendValidationErrors(context, validationErrors)
		return
	}
	// Create person
	person := models.Person{Name: body.Name}

	// pass pointer of data to Create
	result := initializers.DB.Create(&person)

	if result.Error != nil {
		utils.SendErrorResponse(context, result.Error.Error(), 400)
		return
	}
	utils.SendSuccessResponse(context, "Person created successfully", person)
}

// IndexPerson
// @Tags Person CRUD
// @Description get list of all person resource
// @Summary List Person
// @ID index-person
// @Produce json
// @Param name query string false "Name to filter person by"
// @Success 200 {object} utils.JSONSuccessResult{data=[]PersonDto,message}
// @Failure 500 {object} utils.JSONErrorResult{message}
// @Router /person [get]
func IndexPerson(context *gin.Context) {

	var queryParams FilterPersonQueryParams
	err := context.ShouldBindQuery(&queryParams)
	if err != nil {
		utils.SendErrorResponse(context, err.Error(), 400)
		return
	}

	var persons []models.Person
	if len(queryParams.Name) > 0 {
		// Build the query with the dynamic name parameter
		query := "%" + queryParams.Name + "%"
		initializers.DB.Where("name ILIKE ?", query).Find(&persons)
	} else {
		initializers.DB.Find(&persons)
	}

	utils.SendSuccessResponse(context, "All persons retrieved", persons)
}

// ShowPerson
// @Tags Person CRUD
// @Summary Show Person
// @Description Find person resource by id
// @ID show-person
// @Produce json
// @Param id path string true "person ID"
// @Success 200 {object} utils.JSONSuccessResult{data=PersonDto,message}
// @Failure 404 {object} utils.JSONErrorResult{message}
// @Router /person/{id} [get]
func ShowPerson(context *gin.Context) {

	//get ID from url
	id := context.Param("id")

	person, err := getPersonById(id)
	if err == true {
		utils.SendNotFoundErrorResponse(context, "Person could not be found")
		return
	}

	utils.SendSuccessResponse(context, "Person retrieved", person)
}

func getPersonById(id string) (models.Person, bool) {

	//get ID from url
	var person models.Person

	result := initializers.DB.First(&person, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return person, true
	}

	return person, false
}

// UpdatePerson
// @Tags Person CRUD
// @Summary Update Person
// @Description Update person resource by id
// @ID show-person
// @Accept json
// @Produce json
// @Param id path string true "person ID"
// @Param data body UpdatePersonBody true "Person object to update"
// @Success 200 {object} utils.JSONSuccessResult{data=PersonDto,message}
// @Failure 404 {object} utils.JSONErrorResult{message}
// @Failure 422 {array} utils.ValidationError
// @Router /person/{id} [get]
func UpdatePerson(context *gin.Context) {
	var body UpdatePersonBody
	//get ID from url
	id := context.Param("id")
	person, err := getPersonById(id)
	if err == true {
		utils.SendNotFoundErrorResponse(context, "Person could not be found")
		return
	}
	// Get data of request body

	validationErrors := utils.ValidateRequest(context, &body)
	if len(validationErrors) > 0 {
		utils.SendValidationErrors(context, validationErrors)
		return
	}

	//find the person to update
	result := initializers.DB.First(&person, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		utils.SendNotFoundErrorResponse(context, "Person could not be found")
		return
	}

	initializers.DB.Model(&person).Updates(models.Person{
		Name: body.Name,
	})

	utils.SendSuccessResponse(context, "Person retrieved", person)
}

// DeletePerson
// @Tags Person CRUD
// @Summary Delete Person
// @Description Delete Person resource by id
// @ID delete-person
// @Produce json
// @Param id path string true "person ID"
// @Success 200 {object}  utils.JSONSuccessResult{data,message}
// @Failure 404 {object} utils.JSONErrorResult{message}
// @Router /person/{id} [delete]
func DeletePerson(context *gin.Context) {

	//get ID from url
	id := context.Param("id")

	person, err := getPersonById(id)
	if err == true {
		utils.SendNotFoundErrorResponse(context, "Person could not be found")
		return
	}

	initializers.DB.Unscoped().Delete(&models.Person{}, person.ID)

	utils.SendSuccessResponse(context, "Person deleted", nil)
}
