package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/send", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "E-mail sent successfuly",
		})
	})
	r.Run(":80") // listen and serve on 0.0.0.0:8080
}