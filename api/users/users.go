package users

import (
	"log"
	"net/http"
	"stealthy-ninjas/lightning-cards/db"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Service struct {
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) RegisterHandlers(router *gin.Engine) {
	router.GET("/users/create", s.createUUID)
}

func (s *Service) createUUID(gc *gin.Context) {
	res := make(map[string]string)
	res["UUID"] = uuid.New().String()
	db := db.GetService().Db
	rows, err := db.Query("SELECT * FROM cards")
	if err != nil {
		log.Fatal(err)
	}
	var cardV string
	for rows.Next() {
		rows.Scan(&cardV)
		println("Hi", cardV)
	}
	gc.IndentedJSON(http.StatusOK, res)
}
