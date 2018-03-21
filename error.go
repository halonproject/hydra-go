package hydra

import "fmt"

const HYDRA_RESPONSE_TIMEOUT_ERROR = iota + 1

type Error struct {
	code   int
	reason string
}

func newError(code int) Error {
	return Error{
		code:   code,
		reason: "",
	}
}

func (e Error) String() string {
	return e.reason
}

func (e Error) Error() string {
	return fmt.Sprintf("error code %d: %s", e.code, e.reason)
}

func (e Error) Code() int {
	return e.code
}
