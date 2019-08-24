package guesss

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var guess int
var answer int

func TestIsGuessCorrect(t *testing.T) {
	guess = 4500
	answer = 4500
	assert.Equal(t, "You found the answer.", IsGuessCorrect(guess, answer))
	guess = 4000
	assert.Equal(t, "Wrong answer. Your guess is lower.", IsGuessCorrect(guess, answer))
	guess = 4700
	assert.Equal(t, "Wrong answer. Your guess is higher.", IsGuessCorrect(guess, answer))
}
