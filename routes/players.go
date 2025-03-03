package routes

import (
	"golearn/first-api/logger"
	"golearn/first-api/model/player"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func getPlayer(c *gin.Context) {
	playerID, err := strconv.ParseInt(c.Param("playerID"), 10, 64)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		logger.E(err)
		return
	}

	player, err := player.Get(playerID)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		logger.E(err)
		return
	}

	c.JSON(http.StatusOK, player)
}

func getPlayersOfMatch(c *gin.Context) {
	matchID, err := strconv.ParseInt(c.Param("matchID"), 10, 64)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		logger.E(err)
		return
	}

	players, err := player.GetByMatchID(matchID)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		logger.E(err)
		return
	}

	c.JSON(http.StatusOK, players)
}

func postPlayer(c *gin.Context) {
	var player player.Player

	err := c.ShouldBindJSON(&player)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	userID := c.GetInt64("userID")
	matchID, err := strconv.ParseInt(c.Param("matchID"), 10, 64)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		logger.E(err)
		return
	}

	err = player.Save(userID, matchID)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		logger.E(err)
		return
	}

	c.Status(http.StatusCreated)
}

func putPlayer(c *gin.Context) {
	var reqPlayer player.Player
	playerID, err := strconv.ParseInt(c.Param("playerID"), 10, 64)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		logger.E(err)
		return
	}

	err = c.ShouldBindJSON(&reqPlayer)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	dbPlayer, err := player.Get(playerID)
	if dbPlayer == nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "The 'playerID' parameter could not be resolved to a resource."})
		return
	}
	if err != nil {
		c.Status(http.StatusInternalServerError)
		logger.E(err)
		return
	}

	err = dbPlayer.UpdateWith(reqPlayer)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		logger.E(err)
		return
	}

	c.Status(http.StatusOK)
}

func deletePlayer(c *gin.Context) {
	playerID, err := strconv.ParseInt(c.Param("playerID"), 10, 64)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		logger.E(err)
		return
	}

	player, err := player.Get(playerID)
	if player == nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "The 'playerID' parameter could not be resolved to a resource."})
		return
	}
	if err != nil {
		c.Status(http.StatusInternalServerError)
		logger.E(err)
		return
	}

	err = player.Delete()
	if err != nil {
		c.Status(http.StatusInternalServerError)
		logger.E(err)
		return
	}

	c.Status(http.StatusOK)
}
