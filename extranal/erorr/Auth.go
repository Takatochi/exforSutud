package erorr

import "fmt"

type AuthError struct {
	Code int
	Text string
}

func NewAuthError(code int, text string) error {
	return &AuthError{code, text}
}
func (e *AuthError) Error() string {
	return fmt.Sprintf("Code\t%d, message\t%s", e.Code, e.Text)
}

var (
	ErrInvalidCredentials = &AuthError{Code: 200, Text: "Hello World"}
)
