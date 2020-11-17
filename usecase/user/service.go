package user

import (
	"encoding/json"
	"net/http"

	"com.go-crud/entity"
	"com.go-crud/infrastructure"
	"github.com/globalsign/mgo"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

var dao infrastructure.UserDAO

// Init sets initial setup for user service with db connection
func Init(db *mgo.Database) {
	dao = infrastructure.UserDAO{Collection: db.C("users")}
}

// GetAll get all users
func GetAll(w http.ResponseWriter, r *http.Request) {
	users, err := dao.GetAll()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, users)
}

// GetByID get unique user by ID
func GetByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	movie, err := dao.GetByID(params["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Movie ID")
		return
	}

	respondWithJSON(w, http.StatusOK, movie)
}

// Create creates a new user
func Create(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var user entity.User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		log.Errorln("Error on decoder:", err)
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	userReady, err := user.ReadyToCreate()
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := dao.Create(userReady); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, userReady)
}

// Update updates an user
func Update(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	params := mux.Vars(r)

	var user entity.User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	userReady, err := user.ReadyToUpdate()
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := dao.Update(params["id"], userReady); err != nil {
		log.Error(params["id"])
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, map[string]string{"result": userReady.Name + " successfully updated!"})
}

// Delete deletes an user
func Delete(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	params := mux.Vars(r)

	if err := dao.Delete(params["id"]); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	log.Error(msg)
	respondWithJSON(w, code, map[string]string{"error": msg})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
