package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"errors"

)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		
		var err error
		var errCode error
		var resp *http.Response

		resp, err = http.Get("http://audit/log")
		if resp.StatusCode != 200 {
			errCode = errors.New("Error calling service")
		}
		resp, err = http.Get("http://coverage/check")
		if resp.StatusCode != 200 {
			errCode = errors.New("Error calling service")
		}
		resp, err = http.Get("http://mailer/send")
		if resp.StatusCode != 200 {
			errCode = errors.New("Error calling service")
		}

		if err != nil {
			c.JSON(500, gin.H{"error": err})
			return
		}

		if errCode != nil {
			c.JSON(500, gin.H{"error": errCode})
			return
		}

		c.JSON(200, gin.H{
			"message": "ok!",
		})
	})
	r.Run(":80") // listen and serve on 0.0.0.0:8080
}