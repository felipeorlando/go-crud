package user

import (
	"errors"

	"com.go-crud/entity"
)

func validateExistingOfAllFields(u entity.UserSchema) error {
	if u.Name == "" || u.Age == 0 || u.Email == "" || u.Password == "" || u.Address == "" {
		return errors.New("You need to pass all params")
	}

	return nil
}

func validateEmptynessOfAllFields(u entity.UserSchema) error {
	if u.Name == "" && u.Age == 0 && u.Email == "" && u.Password == "" && u.Address == "" {
		return errors.New("You need to pass at least one param")
	}

	return nil
}
