package errors

// Error in server
type Error struct {
	Code    int
	Message string
}

func (e Error) Error() string {
	return e.Message
}

// NewStringError new a error with message
func NewStringError(msg string, code int) error {
	return Error{
		Code:    code,
		Message: msg,
	}
}
