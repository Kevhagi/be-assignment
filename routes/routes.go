package routes

import "github.com/gin-gonic/gin"

func RouteInit(g *gin.RouterGroup) {
	UserRoutes(g)
	TransactionRoutes(g)
}
