package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func testGuesss(t *testing.T) {
	guess := 4500
	answer := 4500
	assert.True(t, isAnswerCorrect(guess, answer))
}
