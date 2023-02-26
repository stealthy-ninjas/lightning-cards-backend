package main

import (
	"net/http"
	"stealthy-ninjas/lightning-cards/api/games"
	"stealthy-ninjas/lightning-cards/api/players"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://127.0.0.1:5173"},
	}))

	// handlers
	gameService := games.NewService()
	playerService := players.NewService()
	gameService.RegisterHandlers(router)
	playerService.RegisterHandlers(router)
	router.GET("/health", healthCheck)

	router.Run("localhost:8080")
}

func healthCheck(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, "Health ok")
}
