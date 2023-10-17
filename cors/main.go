package main

import (
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	config := cors.Config{
		AllowOrigins: []string{
			"http://localhost:9090",
		},
		AllowMethods: []string{"GET"},
		AllowHeaders: []string{"Origin"},
		ExposeHeaders: []string{
			"Content-Length",
		},
		// Set cache, next time preflight
		// is performed. Update this to a higher
		// time to prevent frequent preflight requests
		MaxAge: 1 * time.Second,
	}

	// cm is short for CORS middleware
	cm := cors.New(config)

	r.Use(cm)
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:x
}
