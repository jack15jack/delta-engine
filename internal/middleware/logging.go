package middleware

import (
	"bytes"
	"io"
	"log"

	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {

		if c.Request.Body != nil {

			body, _ := io.ReadAll(c.Request.Body)

			c.Request.Body = io.NopCloser(bytes.NewBuffer(body))

			log.Println("REQUEST:", string(body))
		}

		c.Next()
	}
}
