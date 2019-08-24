package guesss

func IsGuessCorrect(guess, answer int) string {
	return NewResponse(guess, answer).GetResponse()
}
