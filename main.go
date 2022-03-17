package main

import (
	"fmt"
	"net/http"

	"github.com/SeongHyunHan/Tugether/db"
	"github.com/SeongHyunHan/Tugether/handlers"
	"github.com/gorilla/mux"
)

func main() {
	db := db.Init()
	h := handlers.New(db)

	r := mux.NewRouter()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Welcome to Tugether API")
	})
	r.HandleFunc("/Accounts", h.CreateAccount).Methods(http.MethodPost)

	http.ListenAndServe(":8080", r)
}
