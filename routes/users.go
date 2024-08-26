package routes

import (
	"eventBooking/models"
	"eventBooking/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func signUp(context *gin.Context) {

	var user models.User // creating an empty struct
	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = user.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create user"})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "User created!"})
}

func login(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = user.ValidateCredentials()
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	token, err := utils.GenerateToken(user.Email, user.ID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create user"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "User logged in!", "token": token})
}

func getUsers(context *gin.Context) {
	users, err := models.GetAllUsers()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch users"})
		return
	}
	context.JSON(http.StatusOK, users)
}
