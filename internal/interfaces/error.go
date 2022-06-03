package interfaces

type Error struct {
	Reason string `json:"reason"`
}

func NewError(err error) Error {
	return Error{
		Reason: err.Error(),
	}
}
