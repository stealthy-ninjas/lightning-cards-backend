package players

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"stealthy-ninjas/lightning-cards/db"

	"github.com/gin-gonic/gin"
)

type Service struct {
	db *sql.DB
}

func NewService() *Service {
	return &Service{
		db: db.GetService().Db,
	}
}

func (s *Service) RegisterHandlers(router *gin.Engine) {
	router.POST("/players/create", s.createUser)
}

func (s *Service) createUser(gc *gin.Context) {
	requestBody := make(map[string]string)
	err := gc.BindJSON(&requestBody)
	if err != nil {
		log.Println(err.Error())
	}

	var uuid string
	err = s.db.QueryRow(
		fmt.Sprintf(
			"INSERT INTO players (username, ready)"+
				"VALUES ('%s', false) RETURNING id",
			requestBody["userName"],
		),
	).Scan(&uuid)
	if err != nil {
		log.Println(err.Error())
	}

	res := make(map[string]string)
	gc.IndentedJSON(http.StatusOK, res)
}
