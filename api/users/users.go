package users

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Service struct {
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) RegisterHandlers(router *gin.Engine) {
	router.GET("/users/random", s.getRandomUser)
}

func (s *Service) getRandomUser(gc *gin.Context) {
	rand.Seed(time.Now().Unix())
	gc.IndentedJSON(http.StatusOK, fmt.Sprint("User", rand.Intn(100)))
}
