package infrastructure

import (
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"

	"com.go-crud/entity"
)

// UserDAO represents DAO of User entity
type UserDAO struct {
	Collection *mgo.Collection
}

// GetAll get all users
func (u *UserDAO) GetAll() ([]entity.User, error) {
	var users []entity.User
	err := u.Collection.Find(bson.M{}).All(&users)
	return users, err
}

// GetByID get unique user by ID
func (u *UserDAO) GetByID(id string) (entity.User, error) {
	var user entity.User
	err := u.Collection.FindId(bson.ObjectIdHex(id)).One(&user)
	return user, err
}

// Create creates a new user
func (u *UserDAO) Create(user entity.User) error {
	err := u.Collection.Insert(&user)
	return err
}

// Delete deletes an user
func (u *UserDAO) Delete(id string) error {
	err := u.Collection.RemoveId(bson.ObjectIdHex(id))
	return err
}

// Update updates an user
func (u *UserDAO) Update(id string, user entity.User) error {
	err := u.Collection.UpdateId(bson.ObjectIdHex(id), &user)
	return err
}
