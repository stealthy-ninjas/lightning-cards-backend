package games

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"stealthy-ninjas/lightning-cards/models"
	"time"

	"github.com/gin-gonic/gin"
)

type Service struct {
	rooms models.Rooms
}

func NewService(rooms models.Rooms) *Service {
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
	rand.Seed(time.Now().UnixMilli())
	roomId := fmt.Sprint(rand.Int())
	fmt.Println("roomId: ", roomId)

	jsonBuf := map[string]string{}
	jsonBytes, _ := ioutil.ReadAll(gc.Request.Body)
	fmt.Println(jsonBytes)
	json.Unmarshal(jsonBytes, &jsonBuf)
	fmt.Println("henlo", jsonBuf)

	// jsonDecoder := json.NewDecoder(gc.Request.Body)
	// jsonDecoder.Decode(&jsonBuf)

	// todo(): player uuid from database comes below
	s.rooms[roomId] = &models.Room{
		Players: make(map[string]models.Player),
	}
	// s.rooms[roomId].Players[jsonBuf["username"]] = models.Player{}
	gc.IndentedJSON(201, map[string]string{"roomId": roomId})
}
