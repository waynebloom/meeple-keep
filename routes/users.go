package routes

import (
	"golearn/first-api/model"
	"golearn/first-api/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func getUsers(c *gin.Context) {
	users, err := model.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "An unspecified error occurred."})
		return
	}

	c.JSON(http.StatusOK, users)
}

func postUser(c *gin.Context) {
	var user model.User

	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
		return
	}

	err = user.Save()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "An unspecified error occurred."})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "the user was successfully created", "user": user})
}

func login(c *gin.Context) {
	var user model.User

	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
		return
	}

	err = user.Validate()
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		return
	}

	token, err := utils.GenerateToken(user.Email, user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not generate a JWT for the session."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "the login was successful.", "token": token})
}
