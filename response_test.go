package guesss

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestResponse(t *testing.T) {
	answer := 4500

	guess := 4000
	assert.Equal(t, "Wrong answer. Your guess is lower.", NewResponse(guess, answer).GetResponse())
	guess = 4700
	assert.Equal(t, "Wrong answer. Your guess is higher.", NewResponse(guess, answer).GetResponse())
	guess = 4500
	assert.Equal(t, "You found the answer.", NewResponse(guess, answer).GetResponse())
}
