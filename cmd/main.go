package main

import (
	"errors"
	"fmt"
	"hash/extranal/erorr"
	"hash/internal/app"
)

func main() {

	if err := app.Run(); err != nil {
		var authError *erorr.AuthError
		var appError *erorr.AppError

		switch {
		case errors.As(err, &authError):
			fmt.Println(authError)
		case errors.As(err, &appError):
			fmt.Println(appError)
		default:
			fmt.Println(err.Error())
		}

	}

}
