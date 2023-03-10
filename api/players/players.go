package players

import (
	"io/ioutil"
	"net/http"
	"stealthy-ninjas/lightning-cards/db"

	"github.com/gin-gonic/gin"
)

type Service struct {
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) RegisterHandlers(router *gin.Engine) {
	router.GET("/players/create", s.createUser)
}

func (s *Service) createUser(gc *gin.Context) {
	res := make(map[string]string)
	bodyAsBytes, err := ioutil.ReadAll(gc.Request.Body)
	jsonBody := string(bodyAsBytes)
	println(jsonBody)
	if err != nil {
		gc.IndentedJSON(http.StatusBadRequest, map[string]string{"message": "could not process request body"})
	}

	db := db.GetService().Db
	result := db.QueryRow("INSERT INTO players (username, ready) VALUES ('eric', false) RETURNING id")

	var uuid string
	err = result.Scan(&uuid)
	if err != nil {
		println(err.Error())
	}
	println(uuid)
	gc.IndentedJSON(http.StatusOK, res)
}
