package main

var guess int
var answer int

func main() {
}

func isGuessCorrect(guess, answer int) string {
	return responceFactory(guess, answer).getResponce()
}
