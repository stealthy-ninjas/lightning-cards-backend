package games

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Service struct {
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) RegisterHandlers(router *gin.Engine) {
	router.GET("/games", s.Get)
}

func (s *Service) Get(gc *gin.Context) {
	gc.IndentedJSON(http.StatusOK, "list of games here")
}
