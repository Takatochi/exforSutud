package app

import (
	"errors"
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

	//if err := r.Run(); err != nil {
	//	return err
	//}
	return erorr.NewAppError(errors.New("gg"), 200)
}
