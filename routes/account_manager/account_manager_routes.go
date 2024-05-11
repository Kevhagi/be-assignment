package accountmanagerroutes

import (
	"be-assignment/prisma/db"

	"github.com/gin-gonic/gin"
)

func AccountManagerRoutes(g *gin.Engine, gr *gin.RouterGroup, db *db.PrismaClient) {
	UserRoutes(g, gr.Group("/user"), db)
}
