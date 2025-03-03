package routes

import (
	"golearn/first-api/logger"
	"golearn/first-api/model/user"
	"golearn/first-api/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func signup(c *gin.Context) {
	var user user.User

	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	err = user.Save()
	if err != nil {
		c.Status(http.StatusInternalServerError)
		logger.E(err)
		return
	}

	c.Status(http.StatusCreated)
}

func login(c *gin.Context) {
	var user user.User

	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	err = user.Validate()
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		return
	}

	exp, token, err := utils.GenerateToken(user.Email, user.ID)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		logger.E(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token, "exp": exp})
}

// this is a utility function for testing purposes
func getUsers(c *gin.Context) {
	users, err := user.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "An unspecified error occurred."})
		return
	}

	c.JSON(http.StatusOK, users)
}
