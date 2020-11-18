package user

import (
	"encoding/json"
	"net/http"

	"com.go-crud/entity"
	"com.go-crud/utils"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/gorilla/mux"
)

var repo *entity.UserRepo

// NewRoutes sets initial setup and returns the resources route
func NewRoutes(db *mgo.Database, router *mux.Router) {
	repo = &entity.UserRepo{Collection: db.C("users")}

	userRouter := router.PathPrefix("/users").Subrouter()
	userRouter.HandleFunc("", GetAll).Methods("GET")
	userRouter.HandleFunc("/{id}", GetByID).Methods("GET")
	userRouter.HandleFunc("", Create).Methods("POST")
	userRouter.HandleFunc("/{id}", Update).Methods("PUT")
	userRouter.HandleFunc("/{id}", Delete).Methods("DELETE")
}

// GetAll get all users
func GetAll(w http.ResponseWriter, r *http.Request) {
	users, err := repo.GetAll()
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, utils.ErrInternalServer.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, users)
}

// GetByID get unique user by ID
func GetByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	movie, err := repo.GetByID(params["id"])
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, utils.ErrBadRequest.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, movie)
}

// Create creates a new user
func Create(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var user entity.UserSchema

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, utils.ErrBadRequest.Error())
		return
	}

	if err := validateExistingOfAllFields(user); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, utils.ErrBadRequest.Error())
		return
	}

	if err := repo.Create(user); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusCreated, user)
}

// Update updates an user
func Update(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	params := mux.Vars(r)

	var u bson.M

	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, utils.ErrBadRequest.Error())
		return
	}

	if err := validateEmptynessOfAllFields(u); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, utils.ErrBadRequest.Error())
		return
	}

	idParam := params["id"]

	if err := repo.Update(idParam, u); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, map[string]string{"result": "User successfully updated!"})
}

// Delete deletes an user
func Delete(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	params := mux.Vars(r)

	if err := repo.Delete(params["id"]); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}
