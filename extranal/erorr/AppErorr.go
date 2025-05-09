package erorr

import "fmt"

type AppError struct {
	Code    int
	Message string
}

func NewAppError(err error, code int) error {
	return &AppError{code, err.Error()}
}

func (e *AppError) Error() string {
	return fmt.Sprintf("Code\t%d, message\t%s", e.Code, e.Message)
}
