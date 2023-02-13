package main

import (
	"net/http"
	"stealthy-ninjas/lightning-cards/api/games"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// handlers
	gamesService := games.NewService()
	gamesService.RegisterHandlers(router)
	router.GET("/health", healthCheck)

	router.Run("localhost:4200")
}

func healthCheck(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, "Health ok")
}
