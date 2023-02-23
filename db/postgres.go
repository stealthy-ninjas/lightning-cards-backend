package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

type Service struct {
	Db *sql.DB
}

func GetService() *Service {
	return service
}

var service *Service

func init() {
	user := os.Getenv("POSTGRES_USER")
	passw := os.Getenv("POSTGRES_PASSWORD")
	dbname := os.Getenv("POSTGRES_DB")
	connStr := fmt.Sprintf("postgresql://%s:%s@db:5432/%s?sslmode=disable", user, passw, dbname)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	service = &Service{
		Db: db,
	}

}

// func NewService() *Service {
// return &Service{}
// }
