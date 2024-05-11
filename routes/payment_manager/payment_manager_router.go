package paymentmanagerroutes

import (
	"be-assignment/prisma/db"

	"github.com/gin-gonic/gin"
)

func PaymentManagerRoutes(g *gin.Engine, gr *gin.RouterGroup, db *db.PrismaClient) {
	TransactionRoutes(g, gr.Group("/transaction"), db)
}
