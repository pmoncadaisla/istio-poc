package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func main() {
	r := gin.Default()
	r.GET("/check", func(c *gin.Context) {
		time.Sleep(5 * time.Second)
		_, err := http.Get("http://iam/check")
		if err != nil {
			c.JSON(500, gin.H{
				"error": err,
			})
		}
		c.JSON(200, gin.H{
			"message": "Coverage found for address",
		})
	})
	r.Run(":80") // listen and serve on 0.0.0.0:8080
}