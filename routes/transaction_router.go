package routes

import (
	"github.com/gin-gonic/gin"
)

func TransactionRoutes(g *gin.RouterGroup) {
	transaction := g.Group("/transaction")

	transaction.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hellow Transaction Test",
		})
	})
}
