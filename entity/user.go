package entity

import (
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"golang.org/x/crypto/bcrypt"
)

// UserSchema entity that represents all the users values on collection
type UserSchema struct {
	ID       bson.ObjectId `bson:"_id" json:"id"`
	Name     string        `bson:"name" json:"name"`
	Age      int           `bson:"age" json:"age"`
	Email    string        `bson:"email" json:"email"`
	Password string        `bson:"password" json:"password"`
	Address  string        `bson:"address" json:"address"`
}

// UserRepo is a struct that represents the users repo on Mongo
type UserRepo struct {
	Collection *mgo.Collection
}

// GetAll get all users
func (r *UserRepo) GetAll() ([]UserSchema, error) {
	var users []UserSchema
	err := r.Collection.Find(bson.M{}).All(&users)
	return users, err
}

// GetByID get unique user by ID
func (r *UserRepo) GetByID(id string) (UserSchema, error) {
	var user UserSchema
	err := r.Collection.FindId(bson.ObjectIdHex(id)).One(&user)
	return user, err
}

// Create creates a new user
func (r *UserRepo) Create(user UserSchema) error {
	pwd, err := generatePassword(user.Password)
	if err != nil {
		return err
	}

	user.ID = bson.NewObjectId()
	user.Password = pwd

	err = r.Collection.Insert(&user)
	return err
}

// Delete deletes an user
func (r *UserRepo) Delete(id string) error {
	err := r.Collection.RemoveId(bson.ObjectIdHex(id))
	return err
}

// Update updates an user
func (r *UserRepo) Update(id string, user UserSchema) error {
	err := r.Collection.UpdateId(bson.ObjectIdHex(id), &user)
	return err
}

func generatePassword(raw string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(raw), 10)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}
