package token

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInitToken(t *testing.T) {
	var secret = "abc"
	var expireTime = 10
	err := InitToken(secret, expireTime)

	assert.NoError(t, err)
}
