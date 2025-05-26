package routes

import (
	"fmt"
	"net/http"

	"example.com/go_events_api/pkgs/models"
	"example.com/go_events_api/pkgs/utils"
	"github.com/gin-gonic/gin"
)

func signup(context *gin.Context) {
	var user models.User

	// 'ShouldBindJSON' method will map the request body properties with the event struct
	// if in the struct property is marked with struct tags - `bindings:"required"`
	// then it will through error when this values not available
	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Not able to parse the data"})
		return
	}

	err = user.Save()
	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Not able to create user now, try again later..."})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "User created !", "user": user})
}

func login(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Not able to parse the data"})
		return
	}

	err = user.ValidateCredentials()
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		return
	}

	token, err := utils.GenerateToken(user.Email, user.ID)

	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Login successful", "token": token})

}
