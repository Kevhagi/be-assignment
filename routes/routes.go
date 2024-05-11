package routes

import (
	"be-assignment/prisma/db"
	accountmanager "be-assignment/routes/account_manager"
	paymentmanager "be-assignment/routes/payment_manager"

	"github.com/gin-gonic/gin"
)

func RouteInit(g *gin.Engine, db *db.PrismaClient) {
	accountmanager.AccountManagerRoutes(g, g.Group("/account-manager"), db)
	paymentmanager.PaymentManagerRoutes(g, g.Group("/payment-manager"), db)
}
