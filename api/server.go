package api

import (
	"net/http"

	"com.go-crud/config"
	"com.go-crud/database"
	"com.go-crud/usecase/user"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

// Start the server
func Start() {
	db := database.ConnectDB(config.DbNameDev, config.DbURIDev)
	defer database.CloseDB()

	r := mux.NewRouter()
	r.MethodNotAllowedHandler = http.HandlerFunc(MethodNotAllowedHandler)
	r.NotFoundHandler = http.HandlerFunc(NotFoundHandler)

	user.NewRoutes(db, r)

	port := config.APIPort
	log.Println("Server running on", port)
	log.Fatalln(http.ListenAndServe(port, r))
}
