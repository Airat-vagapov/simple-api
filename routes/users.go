package routes

import (
	"net/http"

	"example.com/simple-api/models"
	"example.com/simple-api/utils"
	"github.com/gin-gonic/gin"
)

func signup(context *gin.Context) {
	var user models.User

	err := context.BindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse data"})
	}

	err = user.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not save user"})
	}

	context.JSON(http.StatusCreated, gin.H{"message": "User created"})
}

func login(context *gin.Context) {
	var user models.User

	err := context.BindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse data"})
	}

	err = user.ValidateCredentials()

	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		return
	}

	token, err := utils.GenerateToken(user.Email, user.ID)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Auht success", "token": token})
}

func getUsers(context *gin.Context) {
	users, err := models.GetAllUsers()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch data"})
	}

	context.JSON(http.StatusOK, users)
}
