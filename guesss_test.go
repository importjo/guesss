package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsGuessCorrect(t *testing.T) {
	guess = 4500
	answer = 4500
	assert.Equal(t, "You found the answer.", isGuessCorrect(guess, answer))
	guess = 4000
	assert.Equal(t, "Wrong answer.", isGuessCorrect(guess, answer))
	guess = 4700
	assert.Equal(t, "Wrong answer.", isGuessCorrect(guess, answer))
}
