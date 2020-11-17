package entity

import (
	"errors"

	"github.com/globalsign/mgo/bson"
	"golang.org/x/crypto/bcrypt"
)

// User entity
type User struct {
	ID       bson.ObjectId `bson:"_id" json:"id"`
	Name     string        `bson:"name" json:"name"`
	Age      int           `bson:"age" json:"age"`
	Email    string        `bson:"email" json:"email"`
	Password string        `bson:"password" json:"password"`
	Address  string        `bson:"address" json:"address"`
}

// ReadyToCreate returns User ready to be created
func (u User) ReadyToCreate() (User, error) {
	if err := u.validateAll(); err != nil {
		return u, err
	}

	pwd, err := generatePassword(u.Password)
	if err != nil {
		return u, err
	}

	u.ID = bson.NewObjectId()
	u.Password = pwd

	return u, nil
}

// ReadyToUpdate returns User ready to be updated
func (u User) ReadyToUpdate() (User, error) {
	if err := u.validate(); err != nil {
		return u, err
	}

	pwd, err := generatePassword(u.Password)
	if err != nil {
		return u, err
	}

	u.Password = pwd

	return u, nil
}

func generatePassword(raw string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(raw), 10)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func (u User) validateAll() error {
	if u.Name == "" || u.Age == 0 || u.Email == "" || u.Password == "" || u.Address == "" {
		return errors.New("You need to pass all params")
	}

	return nil
}

func (u User) validate() error {
	if u.Name == "" && u.Age == 0 && u.Email == "" && u.Password == "" && u.Address == "" {
		return errors.New("You need to pass at least one param")
	}

	return nil
}
