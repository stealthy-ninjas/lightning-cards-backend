package main

import (
	"net/http"
	"stealthy-ninjas/lightning-cards/api/games"
	"stealthy-ninjas/lightning-cards/api/players"
	"stealthy-ninjas/lightning-cards/sockets"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
	}))

	// handlers
	gameService := games.NewService()
	playerService := players.NewService()
	socketService := sockets.Server{}
	gameService.RegisterHandlers(router)
	playerService.RegisterHandlers(router)
	router.GET("/health", healthCheck)
	router.GET("/ws", socketService.ServeHttp)

	router.Run("localhost:8080")
}

func healthCheck(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, "Health ok")
}
