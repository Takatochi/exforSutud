package app

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"hash/extranal/erorr"
)

type App struct {
}

func Run() error {

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World",
		})
	})
	err := erorr.NewAppError(errors.New("fdfsd"), 2000)
	//if err := r.Run(); err != nil {
	//	return err
	//}
	return fmt.Errorf(" falide Run()  %w", err)
}
