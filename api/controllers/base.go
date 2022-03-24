package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/SeongHyunHan/Tugether/api/models"
	"github.com/gorilla/mux"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Server struct {
	DB     *gorm.DB
	Router *mux.Router
}

func (server *Server) Initialize(DBDriver, DBUser, DBPassword, DBPort, DBHost, DBName string) {
	var err error

	DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", DBHost, DBPort, DBUser, DBName, DBPassword)
	server.DB, err = gorm.Open(postgres.Open(DBURL), &gorm.Config{})
	if err != nil {
		fmt.Printf("Cannot connect to %s database", DBDriver)
		log.Fatal("This is the error:", err)
	} else {
		fmt.Printf("Connected to the %s database", DBDriver)
	}

	server.DB.Debug().AutoMigrate(&models.Account{})

	server.Router = mux.NewRouter()

	server.initializeRoutes()
}

func (server *Server) Run(addr string) {
	fmt.Printf("Listening to Port %s\n", addr)
	log.Fatal(http.ListenAndServe(addr, server.Router))
}
