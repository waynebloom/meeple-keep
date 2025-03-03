package routes

import (
	"golearn/first-api/logger"
	"golearn/first-api/model/score"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func getScore(c *gin.Context) {
	scoreID, err := strconv.ParseInt(c.Param("scoreID"), 10, 64)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		logger.E(err)
		return
	}

	score, err := score.Get(scoreID)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		logger.E(err)
		return
	}

	c.JSON(http.StatusOK, score)
}

func getScoresOfPlayer(c *gin.Context) {
	playerID, err := strconv.ParseInt(c.Param("playerID"), 10, 64)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		logger.E(err)
		return
	}

	scores, err := score.GetByPlayerID(playerID)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		logger.E(err)
		return
	}

	c.JSON(http.StatusOK, scores)
}

func postScore(c *gin.Context) {
	var score score.Score

	err := c.ShouldBindJSON(&score)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	userID := c.GetInt64("userID")
	playerID, err := strconv.ParseInt(c.Param("playerID"), 10, 64)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		logger.E(err)
		return
	}

	err = score.Save(userID, playerID)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		logger.E(err)
		return
	}

	c.Status(http.StatusCreated)
}

func putScore(c *gin.Context) {
	var reqScore score.Score
	scoreID, err := strconv.ParseInt(c.Param("scoreID"), 10, 64)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		logger.E(err)
		return
	}

	err = c.ShouldBindJSON(&reqScore)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	dbScore, err := score.Get(scoreID)
	if dbScore == nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "The 'scoreID' parameter could not be resolved to a resource."})
		return
	}
	if err != nil {
		c.Status(http.StatusInternalServerError)
		logger.E(err)
		return
	}

	err = dbScore.UpdateWith(reqScore)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		logger.E(err)
		return
	}

	c.Status(http.StatusOK)
}

func deleteScore(c *gin.Context) {
	scoreID, err := strconv.ParseInt(c.Param("scoreID"), 10, 64)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		logger.E(err)
		return
	}

	score, err := score.Get(scoreID)
	if score == nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "The 'scoreID' parameter could not be resolved to a resource."})
		return
	}
	if err != nil {
		c.Status(http.StatusInternalServerError)
		logger.E(err)
		return
	}

	err = score.Delete()
	if err != nil {
		c.Status(http.StatusInternalServerError)
		logger.E(err)
		return
	}

	c.Status(http.StatusOK)
}
