package handler

import (
	"API/config"
	"API/schemas"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	logger *config.Logger
	db *gorm.DB
)

func Init() {
	logger = config.GetLogger("handler")
	db = config.GetSQLite()
}

func CreateOpening(ctx *gin.Context) {
	// Validating JSON
	request := OpeningRequest{}
	ctx.BindJSON(&request)

	if err := request.ValidateCreateOpening(); err != nil {
		logger.Errorf("invalid request: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	// Setting opening and inserting it into the database
	opening := schemas.Opening{
		Role: request.Role,
		Company: request.Company,
		Location: request.Location,
		Link: request.Link,
		Remote: *request.Remote,
		Salary: request.Salary,
	}

	if err := db.Create(&opening).Error; err != nil {
		logger.Errorf("error creating opening: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error creating opening on database")
		return
	}
	
	sendSuccess(ctx, "create-opening", opening)
}

func ShowOpening(ctx *gin.Context) {
	// Getting 'id' query parameter 
	id := ctx.Query("id")
	if id == "" {
		logger.Error("query parameter (id) is required")
		sendError(ctx, http.StatusBadRequest, errParamRequired("id", "query-parameter").Error())
		return
	}

	// Searching for the opening
	opening := schemas.Opening{}
	
	if err := db.First(&opening, id).Error; err != nil {
		logger.Errorf("opening of id: %v not found", id)
		sendError(ctx, http.StatusNotFound, fmt.Sprintf("opening with id: %v not found", id))
		return
	}

	sendSuccess(ctx, "show-opening", opening)
}

func DeleteOpening(ctx *gin.Context) {
	// Getting 'id' query parameter
	id := ctx.Query("id")
	if id == "" {
		logger.Error("query parameter (id) is required")
		sendError(ctx, http.StatusBadRequest, errParamRequired("id", "query-parameter").Error())
		return
	}

	// Searching for the opening and deleting it
	opening := schemas.Opening{}

	if err := db.First(&opening, id).Error; err != nil {
		logger.Errorf("id: %v not found", id)
		sendError(ctx, http.StatusNotFound, fmt.Sprintf("opening with id: %v not found", id))
		return
	}

	if err := db.Delete(&opening).Error; err != nil {
		logger.Errorf("error deleting opening with id: %v, %v", id, err)
		sendError(ctx, http.StatusInternalServerError, fmt.Sprintf("error deleting opening with id: %v", id))
		return
	}
	
	sendSuccess(ctx, "delete-opening", opening)
}

func UpdateOpening(ctx *gin.Context) {
	// Validating JSON
	request := OpeningRequest{}
	ctx.BindJSON(&request)

	if err := request.ValidateUpdateOpening(); err != nil {
		logger.Errorf("validation error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
	}

	// Getting 'id' query parameter
	id := ctx.Query("id")
	if id == "" {
		logger.Error("query parameter (id) is required")
		sendError(ctx, http.StatusBadRequest, errParamRequired("id", "query-parameter").Error())
		return
	}

	// Searching for the opening and updating the requested fields
	opening := schemas.Opening{}

	if err := db.First(&opening, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, "opening not found")
		return
	}

	if request.Role != "" {
		opening.Role = request.Role
	}
    if request.Company != "" {
		opening.Company = request.Company
	}
    if request.Location != "" {
        opening.Location = request.Location
    }
    if request.Link != "" {
		opening.Link = request.Link
	}
    if request.Remote != nil {
		opening.Remote = *request.Remote
	}
    if request.Salary > 0 {
		opening.Salary = request.Salary
	}

	if err := db.Save(&opening).Error; err != nil {
		logger.Errorf("error updating opening: %v", err)
		sendError(ctx, http.StatusInternalServerError, "error updating opening")
		return
	}

	sendSuccess(ctx, "update-opening", opening)
}

func ListOpenings(ctx *gin.Context) {
	// Searching for all the openings
	openings := []schemas.Opening{}

	if err := db.Find(&openings).Error; err != nil {
		logger.Error("error listing openings")
		sendError(ctx, http.StatusInternalServerError, "error listing openings")
		return
	}

	sendSuccess(ctx, "list-openings", openings)
}