package games

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"stealthy-ninjas/lightning-cards/db"
	"stealthy-ninjas/lightning-cards/models"

	"github.com/gin-gonic/gin"
)

type Service struct {
	db    *sql.DB
	rooms models.Rooms
}

func NewService(rooms models.Rooms) *Service {
	return &Service{rooms: rooms, db: db.GetService().Db}
}

func (s *Service) RegisterHandlers(router *gin.Engine) {
	router.GET("/games", s.get)
	router.POST("/create", s.createRoom)
}

func (s *Service) get(gc *gin.Context) {
	gc.IndentedJSON(http.StatusOK, "list of games here")
}

func (s *Service) createRoom(gc *gin.Context) {
	requestBody := make(map[string]string)
	err := gc.BindJSON(&requestBody)
	if err != nil {
		log.Println(err.Error())
	}

	// todo(): get uuid of non hardcoded host
	var uuid string
	err = s.db.QueryRow(
		"SELECT id FROM players WHERE username = 'Grater'",
	).Scan(&uuid)
	if err != nil {
		log.Println(err)
	}

	var roomId string
	err = s.db.QueryRow(
		fmt.Sprintf(
			"INSERT INTO rooms (game_status, host) VALUES (false, '%s') RETURNING id", uuid,
		),
	).Scan(&roomId)
	if err != nil {
		log.Println(err)
	}

	jsonBuf := map[string]string{}
	jsonBytes, _ := ioutil.ReadAll(gc.Request.Body)
	json.Unmarshal(jsonBytes, &jsonBuf)

	// jsonDecoder := json.NewDecoder(gc.Request.Body)
	// jsonDecoder.Decode(&jsonBuf)

	// todo(): player uuid from database comes below
	s.rooms[roomId] = &models.Room{
		Players: make(map[string]models.Player),
	}
	// s.rooms[roomId].Players[jsonBuf["username"]] = models.Player{}
	gc.IndentedJSON(201, map[string]string{"roomId": roomId})
}
