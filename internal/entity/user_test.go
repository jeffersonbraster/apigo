package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	user, err := NewUser("Jeje", "jeje@gmail.com", "123456")
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.NotEmpty(t, user.ID)
	assert.NotEmpty(t, "Jeje", user.Name)
	assert.NotEmpty(t, user.Password)
	assert.Equal(t, "Jeje", user.Name)
	assert.Equal(t, "jeje@gmail.com", user.Email)
}

func TestUser_ValidadePassword(t *testing.T) {
	user, err := NewUser("Jeje", "jeje@gmail.com", "123456")
	assert.Nil(t, err)
	assert.True(t, user.ValidatePassword("123456"))
	assert.False(t, user.ValidatePassword("1234567"))
	assert.NotEqual(t, "123456", user.Password)
}