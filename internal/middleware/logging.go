package middleware

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {

		start := time.Now()

		c.Next()

		log.Printf(
			"%s | %d | %s | %v",
			c.Request.Method,
			c.Writer.Status(),
			c.Request.URL.Path,
			time.Since(start),
		)
	}
}
