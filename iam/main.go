package main

import (
	"github.com/gin-gonic/gin"
 	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/check", func(c *gin.Context) {
		_, err := http.Get("http://mailer/send")
		if err != nil {
			c.JSON(500, gin.H{
				"error": err,
			})
		}
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run(":80") // listen and serve on 0.0.0.0:8080
}