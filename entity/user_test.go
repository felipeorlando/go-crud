package entity_test

import (
	"errors"
	"testing"

	"com.go-crud/utils"

	"com.go-crud/config"
	"com.go-crud/database"
	"com.go-crud/entity"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

var db *mgo.Database
var repo *entity.UserRepo

func TestMain(t *testing.M) {
	db = database.ConnectDB(config.DbNameTest, config.DbURITest)
	repo = &entity.UserRepo{Collection: db.C("users")}

	t.Run()

	clearUsersCollection()
	database.CloseDB()
}

func TestGetAllEmptyUsers(t *testing.T) {
	clearUsersCollection()

	users, err := repo.GetAll()
	assert.Nil(t, err)
	assert.Equal(t, users, []entity.UserSchema([]entity.UserSchema(nil)))
	assert.Equal(t, getUsersCount(), 0)
}

func TestGetAllHappyPath(t *testing.T) {
	clearUsersCollection()

	assert.Equal(t, getUsersCount(), 0)

	u := generateFakeUser("lorem@gmail.com")
	insertNewUserToDb(u)

	users, err := repo.GetAll()
	assert.Nil(t, err)
	assert.Equal(t, users[0].Email, u.Email)
	assert.Equal(t, getUsersCount(), 1)
}

func TestGetByIdWrongId(t *testing.T) {
	clearUsersCollection()

	insertNewUserToDb(generateFakeUser("lorem@gmail.com"))

	_, err := repo.GetByID("5fb343d70f8f6d5f01c0d785")

	assert.Equal(t, err, utils.ErrNotFound)
}

func TestGetByIDHappyPath(t *testing.T) {
	clearUsersCollection()

	firstUserEmail := "lorem@gmail.com"
	u, err := insertNewUserToDb(generateFakeUser(firstUserEmail))

	assert.Nil(t, err)

	if err != nil {
		log.Error(errors.New("Error on TestGetByIdHappyPath inserts"))
	}

	user, err := repo.GetByID(u.ID.Hex())

	assert.Equal(t, u.Email, user.Email)
	assert.Nil(t, err)
}

func TestCreateHappyPath(t *testing.T) {
	clearUsersCollection()

	u := generateFakeUser("lorem@gmail.com")

	createdUser, err := repo.Create(u)

	assert.Nil(t, err)
	assert.Equal(t, getUsersCount(), 1)
	assert.Equal(t, createdUser.Email, u.Email)
}

func TestDeleteHappyPath(t *testing.T) {
	clearUsersCollection()

	u := generateFakeUser("lorem@gmail.com")
	createdUser, _ := insertNewUserToDb(u)

	assert.Equal(t, getUsersCount(), 1)

	err := repo.Delete(createdUser.ID.Hex())

	assert.Nil(t, err)
	assert.Equal(t, getUsersCount(), 0)
}

func TestUpdateHappyPath(t *testing.T) {
	clearUsersCollection()

	u := generateFakeUser("lorem@gmail.com")
	createdUser, _ := insertNewUserToDb(u)

	var updatedUser entity.UserSchema
	_, err := repo.Update(createdUser.ID.Hex(), bson.M{"name": "Clayton"})

	assert.Nil(t, err)

	err = repo.Collection.Find(bson.M{"name": "Clayton"}).One(&updatedUser)

	assert.Nil(t, err)
	assert.Equal(t, getUsersCount(), 1)
	assert.Equal(t, updatedUser.Name, "Clayton")
	assert.Equal(t, updatedUser.Email, "lorem@gmail.com")
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
	repo.Collection.DropCollection()
}
