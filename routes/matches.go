package routes

import (
	"golearn/first-api/logger"
	"golearn/first-api/model/match"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func getMatch(c *gin.Context) {
	matchID, err := strconv.ParseInt(c.Param("matchID"), 10, 64)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		logger.E(err)
		return
	}

	match, err := match.Get(matchID)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		logger.E(err)
		return
	}

	c.JSON(http.StatusOK, match)
}

func getMatchesOfGame(c *gin.Context) {
	gameID := c.GetInt64("gameID")
	matches, err := match.GetByGameID(gameID)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		logger.E(err)
		return
	}

	c.JSON(http.StatusOK, matches)
}

func postMatch(c *gin.Context) {
	var match match.Match

	err := c.ShouldBindJSON(&match)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	userID := c.GetInt64("userID")
	gameID := c.GetInt64("gameID")

	err = match.Save(userID, gameID)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		logger.E(err)
		return
	}

	c.Status(http.StatusCreated)
}

func putMatch(c *gin.Context) {
	var reqMatch match.Match
	matchID, err := strconv.ParseInt(c.Param("matchID"), 10, 64)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		logger.E(err)
		return
	}

	err = c.ShouldBindJSON(&reqMatch)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	dbMatch, err := match.Get(matchID)
	if dbMatch == nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "The 'matchID' parameter could not be resolved to a resource."})
		return
	}
	if err != nil {
		c.Status(http.StatusInternalServerError)
		logger.E(err)
		return
	}

	err = dbMatch.UpdateWith(reqMatch)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		logger.E(err)
		return
	}

	c.Status(http.StatusOK)
}

func deleteMatch(c *gin.Context) {
	matchID, err := strconv.ParseInt(c.Param("matchID"), 10, 64)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		logger.E(err)
		return
	}

	match, err := match.Get(matchID)
	if match == nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "The 'matchID' parameter could not be resolved to a resource."})
		return
	}
	if err != nil {
		c.Status(http.StatusInternalServerError)
		logger.E(err)
		return
	}

	err = match.Delete()
	if err != nil {
		c.Status(http.StatusInternalServerError)
		logger.E(err)
		return
	}

	c.Status(http.StatusOK)
}
