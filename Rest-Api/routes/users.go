package routes

import (
	"net/http"

	"example.org/rest-api/models"
	"example.org/rest-api/utils"
	"github.com/gin-gonic/gin"
)

func Login(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "bad request"})
		return
	}
	isValidUser := user.CheckCredentails()

	if !isValidUser {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "could not authenticate the user"})
		return
	}

	token, err := utils.GenerateToken(user.Email, user.Id)
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "could not authenticate the user"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Login successful", "token": token})
}

func SignUp(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "bad request"})
		return
	}

	userSaved, err := user.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not save user"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"User": userSaved})
}
