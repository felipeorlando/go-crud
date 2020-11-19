package user_test

import (
	"net/http"
	"testing"

	"com.go-crud/api"
	"com.go-crud/config"
	"com.go-crud/database"
	"com.go-crud/entity"
	"com.go-crud/usecase/user"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"github.com/steinfletcher/apitest"
	"golang.org/x/crypto/bcrypt"
)

var db *mgo.Database

func TestMain(t *testing.M) {
	db = database.ConnectDB(config.DbNameTest, config.DbURITest)

	t.Run()

	clearUsersCollection()
	database.CloseDB()
}

func TestIndexHappyPath(t *testing.T) {
	clearUsersCollection()

	insertNewUserToDb(generateFakeUser("lorem@gmail.com"))

	r := setupRouter()

	apitest.New().
		Handler(r).
		Get("/users").
		Expect(t).
		Status(http.StatusOK).
		End()
}

func TestShowHappyPath(t *testing.T) {
	clearUsersCollection()

	u, _ := insertNewUserToDb(generateFakeUser("lorem@gmail.com"))

	r := setupRouter()

	apitest.New().
		Handler(r).
		Get("/users/" + u.ID.Hex()).
		Expect(t).
		Status(http.StatusOK).
		End()
}

func TestShowWrongId(t *testing.T) {
	clearUsersCollection()

	insertNewUserToDb(generateFakeUser("lorem@gmail.com"))

	r := setupRouter()

	apitest.New().
		Handler(r).
		Get("/users/5fb343d70f8f6d5f01c0d785").
		Expect(t).
		Status(http.StatusBadRequest).
		End()
}

func TestNewHappyPath(t *testing.T) {
	clearUsersCollection()

	params := `{"name": "John", "age": 99, "email": "lorem@gmail.com", "password": "123456", "address": "Lorem"}`

	r := setupRouter()

	apitest.New().
		Handler(r).
		Post("/users").
		Body(params).
		Expect(t).
		Status(http.StatusCreated).
		End()
}

func TestNewMissingOneField(t *testing.T) {
	clearUsersCollection()

	params := `{"name": "John", "age": 99, "email": "lorem@gmail.com", "password": "123456"}`

	r := setupRouter()

	apitest.New().
		Handler(r).
		Post("/users").
		Body(params).
		Expect(t).
		Status(http.StatusBadRequest).
		End()
}

func TestEditHappyPath(t *testing.T) {
	clearUsersCollection()

	u, _ := insertNewUserToDb(generateFakeUser("lorem@gmail.com"))
	params := `{"email": "dolor@gmail.com"}`

	r := setupRouter()

	apitest.New().
		Handler(r).
		Put("/users/" + u.ID.Hex()).
		Body(params).
		Expect(t).
		Status(http.StatusOK).
		End()
}

func TestEditWrongID(t *testing.T) {
	clearUsersCollection()

	insertNewUserToDb(generateFakeUser("lorem@gmail.com"))
	params := `{"email": "dolor@gmail.com"}`

	r := setupRouter()

	apitest.New().
		Handler(r).
		Put("/users/5fb343d70f8f6d5f01c0d785").
		Body(params).
		Expect(t).
		Status(http.StatusBadRequest).
		End()
}

func TestEditEmptyParams(t *testing.T) {
	clearUsersCollection()

	u, _ := insertNewUserToDb(generateFakeUser("lorem@gmail.com"))

	r := setupRouter()

	apitest.New().
		Handler(r).
		Put("/users/" + u.ID.Hex()).
		Body("{}").
		Expect(t).
		Status(http.StatusBadRequest).
		End()
}

func TestDeleteHappyPath(t *testing.T) {
	clearUsersCollection()

	u, _ := insertNewUserToDb(generateFakeUser("lorem@gmail.com"))
	r := setupRouter()

	apitest.New().
		Handler(r).
		Delete("/users/" + u.ID.Hex()).
		Expect(t).
		Status(http.StatusOK).
		End()
}

func TestDeleteWrongID(t *testing.T) {
	clearUsersCollection()

	insertNewUserToDb(generateFakeUser("lorem@gmail.com"))
	r := setupRouter()

	apitest.New().
		Handler(r).
		Delete("/users/5fb343d70f8f6d5f01c0d785").
		Expect(t).
		Status(http.StatusInternalServerError).
		End()
}

func getUsersCount() int {
	c, err := db.C("users").Count()
	if err != nil {
		log.Error(err)
	}

	return c
}

func generateFakeUser(email string) entity.UserSchema {
	return entity.UserSchema{Name: "John", Age: 99, Email: email, Password: "123456", Address: "Lorem"}
}

func insertNewUserToDb(u entity.UserSchema) (entity.UserSchema, error) {
	pwd, err := bcrypt.GenerateFromPassword([]byte(u.Password), 10)
	if err != nil {
		log.Error("Error generating:", err)
		return u, err
	}

	u.ID = bson.NewObjectId()
	u.Password = string(pwd)

	if err := db.C("users").Insert(&u); err != nil {
		log.Error("Error on creating new user:", err)
		return u, err
	}

	return u, nil
}

func clearUsersCollection() {
	db.C("users").DropCollection()
}

func setupRouter() *mux.Router {
	r := mux.NewRouter()
	r.MethodNotAllowedHandler = http.HandlerFunc(api.MethodNotAllowedHandler)
	r.NotFoundHandler = http.HandlerFunc(api.NotFoundHandler)

	user.NewRoutes(db, r)

	return r
}
