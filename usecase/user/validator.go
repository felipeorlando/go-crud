package user

import (
	"errors"

	"com.go-crud/entity"
	"github.com/globalsign/mgo/bson"
)

func validateExistingOfAllFields(u entity.UserSchema) error {
	if u.Name == "" || u.Age == 0 || u.Email == "" || u.Password == "" || u.Address == "" {
		return errors.New("You need to pass all params")
	}

	return nil
}

func validateEmptynessOfAllFields(u bson.M) error {
	if u["name"] == nil && u["age"] == nil && u["email"] == nil && u["password"] == nil && u["address"] == nil {
		return errors.New("You need to pass at least one param")
	}

	return nil
}
