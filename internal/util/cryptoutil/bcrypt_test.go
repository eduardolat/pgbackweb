package cryptoutil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateBcryptHash(t *testing.T) {
	password := "mysecretpassword"

	hash, err := CreateBcryptHash(password)
	assert.NoError(t, err)
	assert.NotEmpty(t, hash)
}

func TestVerifyBcryptHash(t *testing.T) {
	password := "mysecretpassword"

	hash, err := CreateBcryptHash(password)
	assert.NoError(t, err)
	assert.NotEmpty(t, hash)

	err = VerifyBcryptHash(password, hash)
	assert.NoError(t, err)
}

func TestVerifyBcryptHash_InvalidPassword(t *testing.T) {
	password := "mysecretpassword"
	invalidPassword := "invalidpassword"

	hash, err := CreateBcryptHash(password)
	assert.NoError(t, err)
	assert.NotEmpty(t, hash)

	err = VerifyBcryptHash(invalidPassword, hash)
	assert.Error(t, err)
	assert.Equal(t, "invalid password", err.Error())
}
