package entity

import (
	"com.go-crud/utils"
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

// UserUpdateSchema is the schema to validate fields for update
type UserUpdateSchema struct {
	Name     *string `bson:"name" json:"name,omitempty"`
	Age      *int    `bson:"age" json:"age,omitempty"`
	Email    *string `bson:"email" json:"email,omitempty"`
	Password *string `bson:"password" json:"password,omitempty"`
	Address  *string `bson:"address" json:"address,omitempty"`
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
	if err := r.Collection.FindId(bson.ObjectIdHex(id)).One(&user); err != nil {
		return user, utils.ErrNotFound
	}

	return user, nil
}

// Create creates a new user
func (r *UserRepo) Create(u UserSchema) (UserSchema, error) {
	pwd, err := generatePassword(u.Password)
	if err != nil {
		return u, err
	}

	u.ID = bson.NewObjectId()
	u.Password = pwd

	err = r.Collection.Insert(&u)
	return u, err
}

// Delete deletes an user
func (r *UserRepo) Delete(id string) error {
	if err := r.Collection.RemoveId(bson.ObjectIdHex(id)); err != nil {
		return utils.ErrNotFound
	}

	return nil
}

// Update updates an user
func (r *UserRepo) Update(id string, u bson.M) (bson.M, error) {
	if u["password"] != nil {
		pwd, err := generatePassword(u["password"].(string))
		if err != nil {
			return u, err
		}

		u["password"] = pwd
	}

	err := r.Collection.UpdateId(bson.ObjectIdHex(id), bson.M{"$set": u})
	return u, err
}

func generatePassword(raw string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(raw), 10)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}
