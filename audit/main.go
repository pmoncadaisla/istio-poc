package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
)


var client *redis.Client = redis.NewClient(&redis.Options{
	Addr:     "redis:6379",
	Password: "", // no password set
	DB:       0,  // use default DB
})

func main() {
	r := gin.Default()
	r.GET("/log", func(c *gin.Context) {
		err := client.Set("key", "value", 0).Err()
		if err != nil {
			c.JSON(500, gin.H{
				"error": err,
			})
			return
		}
		c.JSON(200, gin.H{
			"message": "Audit logged",
		})
	})
	r.Run(":80") // listen and serve on 0.0.0.0:8080
}