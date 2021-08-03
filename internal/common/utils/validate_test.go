package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsPrice(t *testing.T) {
	v := NewValidate()

	err := v.Var("0", "is_price")
	assert.NoError(t, err)
	err = v.Var("10", "is_price")
	assert.NoError(t, err)
	err = v.Var("10.0", "is_price")
	assert.NoError(t, err)
	err = v.Var("10.01", "is_price")
	assert.NoError(t, err)

	err = v.Var("10.012", "is_price")
	assert.NotNil(t, err)
	err = v.Var("a10.0123", "is_price")
	assert.NotNil(t, err)

}

func TestIsUsername(t *testing.T) {
	v := NewValidate()

	err := v.Var("username", "is_username")
	assert.NoError(t, err)
	err = v.Var("username123", "is_username")
	assert.NoError(t, err)
	err = v.Var("Username123", "is_username")
	assert.NoError(t, err)

	err = v.Var("username 123", "is_username")
	assert.NotNil(t, err)
	err = v.Var("username1234567890", "is_username")
	assert.NotNil(t, err)
}
