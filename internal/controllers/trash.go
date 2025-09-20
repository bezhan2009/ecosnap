package controllers

import (
	"ecosnap/internal/app/models"
	"ecosnap/internal/app/service"
	"ecosnap/internal/app/service/validators"
	"ecosnap/internal/controllers/middlewares"
	"ecosnap/pkg/errs"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetAllTrashes(c *gin.Context) {
	isValid, month := validators.ValidateMonth(c)
	if !isValid {
		HandleError(c, errs.ErrInvalidMonth)
		return
	}

	isValid, year := validators.ValidateYear(c)
	if !isValid {
		HandleError(c, errs.ErrInvalidYear)
		return
	}

	trashes, err := service.GetAllTrashes(month, year)
	if err != nil {
		HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, trashes)
}

func GetTrashByID(c *gin.Context) {
	trashIdStr := c.Param("id")
	trashId, err := strconv.Atoi(trashIdStr)
	if err != nil {
		HandleError(c, errs.ErrInvalidID)
		return
	}

	trash, err := service.GetTrashByID(trashId)
	if err != nil {
		HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, trash)
}

func CreateTrash(c *gin.Context) {
	userID := c.GetUint(middlewares.UserIDCtx)
	var trash models.Trash

	trash.UserID = userID

	if err := c.ShouldBindJSON(&trash); err != nil {
		HandleError(c, errs.ErrValidationFailed)
		return
	}

	err := service.CreateTrash(trash)
	if err != nil {
		HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Trash created successfully"})
}

func DeleteTrash(c *gin.Context) {
	trashIdStr := c.Param("id")
	trashId, err := strconv.Atoi(trashIdStr)
	if err != nil {
		HandleError(c, errs.ErrInvalidID)
		return
	}

	err = service.DeleteTrash(trashId)
	if err != nil {
		HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Trash deleted successfully"})
}
