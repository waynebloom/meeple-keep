package middleware

import (
	"golearn/first-api/logger"
	"golearn/first-api/model/game"
	"golearn/first-api/utils"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func Authenticate(c *gin.Context) {
	auth := c.Request.Header.Get("Authorization")
	token := strings.Split(auth, "Bearer ")[1]

	if token == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "You need to provide an auth token."})
		return
	}

	userId, err := utils.VerifyToken(token)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		return
	}

	c.Set("userID", userId)
	c.Next()
}

func VerifyGameOwnership(c *gin.Context) {
	userID := c.GetInt64("userID")
	gameID, err := strconv.ParseInt(c.Param("gameID"), 10, 64)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		logger.W(err)
		return
	}

	game, err := game.Get(gameID)
	if game == nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "The 'gameID' parameter could not be resolved to a resource."})
		return
	}
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		logger.E(err)
		return
	}

	if userID != game.OwnerID {
		c.AbortWithStatus(http.StatusForbidden)
		return
	}

	c.Set("gameID", game.ID)
	c.Next()
}
