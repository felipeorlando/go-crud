package user

import (
	"encoding/json"
	"net/http"

	"com.go-crud/entity"
	"github.com/globalsign/mgo"
	"github.com/gorilla/mux"
)

var repo *entity.UserRepo

// Init sets initial setup for user service with db connection
func Init(db *mgo.Database) {
	repo = &entity.UserRepo{Collection: db.C("users")}
}

// GetAll get all users
func GetAll(w http.ResponseWriter, r *http.Request) {
	users, err := repo.GetAll()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, users)
}

// GetByID get unique user by ID
func GetByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	movie, err := repo.GetByID(params["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Movie ID")
		return
	}

	respondWithJSON(w, http.StatusOK, movie)
}

// Create creates a new user
func Create(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var user entity.UserSchema

	if err := validateExistingOfAllFields(user); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	if err := repo.Create(user); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, user)
}

// Update updates an user
func Update(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	params := mux.Vars(r)

	var user entity.UserSchema

	if err := validateEmptynessOfAllFields(user); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	if err := repo.Update(params["id"], user); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, map[string]string{"result": user.Name + " successfully updated!"})
}

// Delete deletes an user
func Delete(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	params := mux.Vars(r)

	if err := repo.Delete(params["id"]); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJSON(w, code, map[string]string{"error": msg})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
