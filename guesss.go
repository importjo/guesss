package main

var guess int
var answer int

func main() {
}

func isGuessCorrect(guess, answer int) string {
	if guess == answer {
		return "You found the answer."
	}

	if guess < answer {
		return "Wrong answer. Your guess is lower."
	}

	return "Wrong answer. Your guess is higher."
}
