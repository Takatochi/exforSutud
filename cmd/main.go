package main

import (
	"fmt"
	"hash/extranal/erorr"
	"hash/internal/app"
	"log"
)

func main() {

	if err := app.Run(); err != nil {
		if e, ok := err.(*erorr.AuthError); ok {
			log.Fatal(fmt.Sprintf("error [%d]: %v", e.Code, "fdfd"))
		} else {
			log.Fatal("error occurred")
		}

	}

}
