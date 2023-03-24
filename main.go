package main

import (
	"net/http"
	"stealthy-ninjas/lightning-cards/api/games"
	"stealthy-ninjas/lightning-cards/api/players"
	"stealthy-ninjas/lightning-cards/models"
	"stealthy-ninjas/lightning-cards/sockets"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	rooms := models.Rooms{}

	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
	}))

	// register handlers
	gameService := games.NewService(rooms)
	gameService.RegisterHandlers(router)
	playerService := players.NewService()
	playerService.RegisterHandlers(router)
	socketService := sockets.NewService(rooms)
	router.GET("/ws", socketService.ServeHttp)

	router.GET("/health", healthCheck)

	router.Run("localhost:8080")
}

func healthCheck(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, "Health ok")
}
