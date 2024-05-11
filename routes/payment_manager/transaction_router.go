package paymentmanagerroutes

import (
	"be-assignment/prisma/db"

	"github.com/gin-gonic/gin"
)

func TransactionRoutes(g *gin.Engine, gr *gin.RouterGroup, db *db.PrismaClient) {
	gr.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hellow Transaction Test",
		})
	})
}
