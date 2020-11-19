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
	userRouter.HandleFunc("", Index).Methods("GET")
	userRouter.HandleFunc("/{id}", Show).Methods("GET")
	userRouter.HandleFunc("", New).Methods("POST")
	userRouter.HandleFunc("/{id}", Edit).Methods("PUT")
	userRouter.HandleFunc("/{id}", Destroy).Methods("DELETE")
}

// Index get all users
func Index(w http.ResponseWriter, r *http.Request) {
	users, err := repo.GetAll()
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, utils.ErrInternalServer.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, users)
}

// Show get unique user by ID
func Show(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	movie, err := repo.GetByID(params["id"])
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, utils.ErrBadRequest.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, movie)
}

// New creates a new user
func New(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var user entity.UserSchema

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, utils.ErrBadRequest.Error())
		return
	}

	if err := ValidateExistingOfAllFields(user); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	savedUser, err := repo.Create(user)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusCreated, savedUser)
}

// Edit updates an user
func Edit(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	params := mux.Vars(r)

	var u bson.M

	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, utils.ErrBadRequest.Error())
		return
	}

	if err := ValidateEmptynessOfAllFields(u); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	idParam := params["id"]

	_, err := repo.Update(idParam, u)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, map[string]string{"result": "User successfully updated!"})
}

// Destroy deletes an user
func Destroy(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	params := mux.Vars(r)

	if err := repo.Delete(params["id"]); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}
