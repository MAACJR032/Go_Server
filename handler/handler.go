package handler

import (
	"API/config"
	"API/schemas"
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
	request := CreateOpeningRequest{}
	ctx.BindJSON(&request)

	err	:= request.Validate()
	if err != nil {
		logger.Errorf("invalid request: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	opening := schemas.Opening{
		Role: request.Role,
		Company: request.Company,
		Location: request.Location,
		Link: request.Link,
		Remote: *request.Remote,
		Salary: request.Salary,
	}

	err = db.Create(&opening).Error
	if err != nil {
		logger.Errorf("error creating opening: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error creating opening on database")
		return
	}
	
	sendSuccess(ctx, "create-opening", opening)
}

func ShowOpening(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "POST Opening",
	})
}

func DeleteOpening(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "DELETE Opening",
	})
}

func UpdateOpening(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "PUT Opening",
	})
}

func ListOpenings(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "GET Openings",
	})
}