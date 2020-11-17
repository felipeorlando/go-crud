package main

import (
	"net/http"

	"com.go-crud/config"
	"com.go-crud/database"
	"com.go-crud/usecase/user"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

func main() {
	db := database.ConnectDB()
	defer database.CloseDB()

	user.Init(db)

	r := mux.NewRouter()
	r.HandleFunc("/api/v1/users", user.GetAll).Methods("GET")
	r.HandleFunc("/api/v1/users/{id}", user.GetByID).Methods("GET")
	r.HandleFunc("/api/v1/users", user.Create).Methods("POST")
	r.HandleFunc("/api/v1/users/{id}", user.Update).Methods("PUT")
	r.HandleFunc("/api/v1/users/{id}", user.Delete).Methods("DELETE")

	port := config.APIPort
	log.Println("Server running on", port)
	log.Fatalln(http.ListenAndServe(port, r))
}
