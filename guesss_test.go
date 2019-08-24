package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGuesss(t *testing.T) {
	guess = 4500
	answer = 4500
	assert.True(t, isAnswerCorrect(guess, answer))
	guess = 4000
	assert.False(t, isAnswerCorrect(guess, answer))
	guess = 4700
	assert.False(t, isAnswerCorrect(guess, answer))
}
