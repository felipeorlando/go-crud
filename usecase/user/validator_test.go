package user_test

import (
	"testing"

	"com.go-crud/utils"

	"com.go-crud/entity"
	"com.go-crud/usecase/user"
	"github.com/globalsign/mgo/bson"
	"github.com/stretchr/testify/assert"
)

func TestValidateExistingOfAllFields(t *testing.T) {
	u := entity.UserSchema{Name: "Felipe"}
	assert.Equal(t, user.ValidateExistingOfAllFields(u), utils.ErrBadRequest)

	u = entity.UserSchema{}
	assert.Equal(t, user.ValidateExistingOfAllFields(u), utils.ErrBadRequest)

	u = entity.UserSchema{Name: "Felipe", Age: 99, Email: "fobsouza@gmail.com", Password: "123", Address: "Lorem"}
	assert.Nil(t, user.ValidateExistingOfAllFields(u))
}

func TestValidateEmptynessOfAllFields(t *testing.T) {
	u := bson.M{"name": "Felipe"}
	assert.Nil(t, user.ValidateEmptynessOfAllFields(u))

	u = bson.M{}
	assert.Equal(t, user.ValidateEmptynessOfAllFields(u), utils.ErrBadRequest)
}
