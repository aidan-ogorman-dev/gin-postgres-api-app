package controller

import (
	"diary_app/helper"
	"diary_app/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddEntry(ctx *gin.Context) {
	var input models.Entry
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	currentUser, err := helper.CurrentUser(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	input.UserID = currentUser.ID
	entry, err := input.Save()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	ctx.JSON(http.StatusCreated, gin.H{"data": entry})
}

func GetAllEntries(ctx *gin.Context) {
	user, err := helper.CurrentUser(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	ctx.JSON(http.StatusOK, gin.H{"data": user.Entries})
}
