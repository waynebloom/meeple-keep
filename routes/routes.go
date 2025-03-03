package routes

import (
	"github.com/gin-gonic/gin"
	"golearn/first-api/middleware"
)

func RegisterRoutes(server *gin.Engine) {
	server.POST("/signup", signup)
	server.POST("/login", login)

	auth := server.Group("/")
	auth.Use(middleware.Authenticate)
	auth.GET("/games", getGames)
	auth.POST("/games", postGame)

	// all app data requires ownership verification
	gameData := auth.Group("/games/:gameID")
	gameData.Use(middleware.VerifyGameOwnership)

	gameData.GET("", getGame)
	gameData.PUT("", putGame)
	gameData.DELETE("", deleteGame)

	gameData.GET("/matches", getMatchesOfGame)
	gameData.POST("/matches", postMatch)
	gameData.GET("/matches/:matchID", getMatch)
	gameData.PUT("/matches/:matchID", putMatch)
	gameData.DELETE("/matches/:matchID", deleteMatch)

	gameData.GET("/categories", getCategoriesOfGame)
	gameData.POST("/categories", postCategory)
	gameData.GET("/categories/:categoryID", getCategory)
	gameData.PUT("/categories/:categoryID", putCategory)
	gameData.DELETE("/categories/:categoryID", deleteCategory)

	gameData.GET("/matches/:matchID/players", getPlayersOfMatch)
	gameData.POST("/matches/:matchID/players", postPlayer)
	gameData.GET("/matches/:matchID/players/:playerID", getPlayer)
	gameData.PUT("/matches/:matchID/players/:playerID", putPlayer)
	gameData.DELETE("/matches/:matchID/players/:playerID", deletePlayer)

	gameData.GET("/matches/:matchID/players/:playerID/scores", getScoresOfPlayer)
	gameData.POST("/matches/:matchID/players/:playerID/scores", postScore)
	gameData.GET("/matches/:matchID/players/:playerID/scores/:scoreID", getScore)
	gameData.PUT("/matches/:matchID/players/:playerID/scores/:scoreID", putScore)
	gameData.DELETE("/matches/:matchID/players/:playerID/scores/:scoreID", deleteScore)
}
