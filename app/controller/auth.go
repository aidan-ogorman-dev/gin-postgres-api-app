package controller

import (
	"diary_app/helper"
	"diary_app/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Register(ctx *gin.Context) {
	var input models.AuthInput

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newUser := &models.User{
		Username: input.Username,
		Password: input.Password,
	}

	user, err := newUser.Save()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"user": user})
}

func Login(ctx *gin.Context) {
	var input models.AuthInput

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := models.FindUserByUsername(input.Username)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	if err = user.ValidatePassword(input.Password); err != nil {
		ctx.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
		return
	}

	token, err := helper.GenerateJWT(user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"jwt": token})
}
