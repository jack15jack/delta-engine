package middleware

import "github.com/gin-gonic/gin"

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {

		// TEMP: fake user (replace later)
		userID := "user-123"

		c.Set("userID", userID)
		c.Next()
	}
}
