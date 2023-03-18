package games

import (
	"encoding/json"
	"fmt"
	"net/http"
	"stealthy-ninjas/lightning-cards/types"

	"github.com/gin-gonic/gin"
)

type Service struct {
	rooms types.Rooms
}

func NewService(rooms types.Rooms) *Service {
	return &Service{rooms: rooms}
}

func (s *Service) RegisterHandlers(router *gin.Engine) {
	router.GET("/games", s.get)
	router.POST("/create", s.createRoom)
}

func (s *Service) get(gc *gin.Context) {
	gc.IndentedJSON(http.StatusOK, "list of games here")
}

func (s *Service) createRoom(gc *gin.Context) {
	// todo(): get room uuid from backend table
	jsonBuf := map[string]string{}
	jsonDecoder := json.NewDecoder(gc.Request.Body)
	jsonDecoder.Decode(&jsonBuf)
	fmt.Println("Hey!", jsonBuf)
	// todo(): player uuid from database comes below
	s.rooms["r1"] = types.Room{
		Players: map[string]types.Player{
			(jsonBuf["userName"]): types.Player{
				Username: jsonBuf["userName"],
			},
		},
	}
	gc.IndentedJSON(201, map[string]string{"roomId": "r1"})
}
