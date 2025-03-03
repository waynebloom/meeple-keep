package routes

import (
	"database/sql"
	"golearn/first-api/db"
	"golearn/first-api/logger"
	"golearn/first-api/model/game"
	"net/http"

	"github.com/gin-gonic/gin"
)

func getGame(c *gin.Context) {
	game, err := game.Get(c.GetInt64("gameID"))
	if err != nil {
		c.Status(http.StatusInternalServerError)
		logger.E(err)
		return
	}

	if game == nil {
		c.Status(http.StatusOK)
		return
	}

	c.JSON(http.StatusOK, game)
}

func getGames(c *gin.Context) {
	userID := c.GetInt64("userID")

	stmt, err := db.DB.Prepare("SELECT * FROM Game WHERE owner_id = ?")
	if err != nil {
		c.Status(http.StatusInternalServerError)
		logger.E(err)
		return
	}
	defer stmt.Close()

	result, err := stmt.Query(userID)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		logger.E(err)
		return
	}

	var games []game.Game
	for result.Next() {
		var game game.Game

		err = result.Scan(&game.ID, &game.OwnerID, &game.Name, &game.Color, &game.ScoringMode)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			logger.E(err)
			return
		}

		games = append(games, game)
	}

	if len(games) == 0 {
		c.Status(http.StatusNoContent)
		return
	}

	c.JSON(http.StatusOK, games)
}

func postGame(c *gin.Context) {
	var game game.Game
	userID := c.GetInt64("userID")

	err := c.ShouldBindJSON(&game)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	err = game.Save(userID)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		logger.E(err)
		return
	}

	c.Status(http.StatusCreated)
}

func putGame(c *gin.Context) {
	var reqGame game.Game
	gameID := c.GetInt64("gameID")

	err := c.ShouldBindJSON(&reqGame)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	stmt, err := db.DB.Prepare("SELECT * FROM Game WHERE id = ?")
	if err != nil {
		c.Status(http.StatusInternalServerError)
		logger.E(err)
		return
	}

	var dbGame game.Game
	err = stmt.QueryRow(gameID).Scan(
		&dbGame.ID,
		&dbGame.OwnerID,
		&dbGame.Name,
		&dbGame.Color,
		&dbGame.ScoringMode,
	)

	if err == sql.ErrNoRows {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Attempted to update a non-existent resource."})
		return
	}
	if err != nil {
		c.Status(http.StatusInternalServerError)
		logger.E(err)
		return
	}

	err = dbGame.UpdateWith(reqGame)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		logger.E(err)
		return
	}

	c.Status(http.StatusOK)
}

func deleteGame(c *gin.Context) {
	game, err := game.Get(c.GetInt64("gameID"))
	if err != nil {
		c.Status(http.StatusInternalServerError)
		logger.E(err)
		return
	}

	err = game.Delete()
	if err != nil {
		c.Status(http.StatusInternalServerError)
		logger.E(err)
		return
	}

	c.Status(http.StatusOK)
}
