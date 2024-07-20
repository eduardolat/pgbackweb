package cryptoutil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateHash(t *testing.T) {
	password := "mysecretpassword"

	hash, err := CreateHash(password)
	assert.NoError(t, err)
	assert.NotEmpty(t, hash)
}

func TestVerifyHash(t *testing.T) {
	password := "mysecretpassword"

	hash, err := CreateHash(password)
	assert.NoError(t, err)
	assert.NotEmpty(t, hash)

	err = VerifyHash(password, hash)
	assert.NoError(t, err)
}

func TestVerifyHash_InvalidPassword(t *testing.T) {
	password := "mysecretpassword"
	invalidPassword := "invalidpassword"

	hash, err := CreateHash(password)
	assert.NoError(t, err)
	assert.NotEmpty(t, hash)

	err = VerifyHash(invalidPassword, hash)
	assert.Error(t, err)
	assert.Equal(t, "invalid password", err.Error())
}
