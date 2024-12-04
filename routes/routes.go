package routes

import (
	"golearn/first-api/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	// for owner of data

	// GET games
	// POST game
	// PUT game
	// DELETE game

	// GET matches of game
	// POST match
	// PUT match
	// DELETE match
	server.GET("/events", getEvents)
}
