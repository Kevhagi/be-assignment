package routes

import (
	"github.com/gin-gonic/gin"
)

func UserRoutes(g *gin.RouterGroup) {
	user := g.Group("/user")

	user.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hellow User Test",
		})
	})
}
