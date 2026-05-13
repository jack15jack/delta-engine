package routes

import "github.com/gin-gonic/gin"

func RegisterAnalyticsRoutes(api *gin.RouterGroup) {
	analytics := api.Group("/analytics")

	analytics.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "analytics ok"})
	})
}
