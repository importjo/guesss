package guesss

type responce interface {
	getResponce() string
}

type correctResponse struct {
	responce string
}

type lowerResponce struct {
	responce string
}

type higherResponce struct {
	responce string
}

type noResponce struct {
	responce string
}

func (r correctResponse) getResponce() string {
	return r.responce
}

func (r lowerResponce) getResponce() string {
	return r.responce
}

func (r higherResponce) getResponce() string {
	return r.responce
}

func (r noResponce) getResponce() string {
	return r.responce
}

func responceFactory(guess, answer int) responce {
	if guess == answer {
		return correctResponse{responce: "You found the answer."}
	}

	if guess < answer {
		return lowerResponce{responce: "Wrong answer. Your guess is lower."}
	}

	if guess > answer {
		return higherResponce{responce: "Wrong answer. Your guess is higher."}
	}

	return noResponce{responce: "No responce."}
}

func IsGuessCorrect(guess, answer int) string {
	return responceFactory(guess, answer).getResponce()
}
