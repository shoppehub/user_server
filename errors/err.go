package errors

import (
	"encoding/json"
)

type CodeError struct {
	Code int
	Msg  string
}

func (e CodeError) Error() string {
	err, _ := json.Marshal(e)
	return string(err)
}

// func New(code int, msg string) *CodeError {
// 	return &CodeError{
// 		Code: code,
// 		Msg:  msg,
// 	}
// }

func IsCodeError(err error) bool {
	_, ok := err.(CodeError)
	return ok
}

// type EmailNotFound struct {
// 	Email string
// }

// func IsEmailNotFound(err error) bool {
// 	_, ok := err.(EmailNotFound)
// 	return ok
// }

// func (err EmailNotFound) Error() string {
// 	return fmt.Sprintf("email is not found [email: %s]", err.Email)
// }
