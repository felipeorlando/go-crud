package user

import (
	"com.go-crud/utils"

	"com.go-crud/entity"
	"github.com/globalsign/mgo/bson"
)

// ValidateExistingOfAllFields validate existing of all fileds
func ValidateExistingOfAllFields(u entity.UserSchema) error {
	if u.Name == "" || u.Age == 0 || u.Email == "" || u.Password == "" || u.Address == "" {
		return utils.ErrBadRequest
	}

	return nil
}

// ValidateEmptynessOfAllFields validate existing of at least one field
func ValidateEmptynessOfAllFields(u bson.M) error {
	if u["name"] == nil && u["age"] == nil && u["email"] == nil && u["password"] == nil && u["address"] == nil {
		return utils.ErrBadRequest
	}

	return nil
}
