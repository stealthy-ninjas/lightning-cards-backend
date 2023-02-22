package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Service struct {
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) RegisterHandlers(router *gin.Engine) {
	router.GET("/users/create", s.getUUID)
}

func (s *Service) getUUID(gc *gin.Context) {
	println("hiii")
	res := make(map[string]string)
	res["UUID"] = uuid.New().String()
	gc.IndentedJSON(http.StatusOK, res)
}
