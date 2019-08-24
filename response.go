package guesss

type Response interface {
	GetResponse() string
}

type responseImpl struct {
	response string
}

func (r responseImpl) GetResponse() string {
	return r.response
}

func NewResponse(guess, answer int) Response {
	if guess == answer {
		return responseImpl{
			response: "You found the answer.",
		}
	}

	if guess < answer {
		return responseImpl{
			response: "Wrong answer. Your guess is lower.",
		}
	}

	if guess > answer {
		return responseImpl{
			response: "Wrong answer. Your guess is higher.",
		}
	}

	return responseImpl{
		response: "No response.",
	}
}
