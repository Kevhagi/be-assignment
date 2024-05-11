package accountmanagerroutes

import (
	controllerImplement "be-assignment/controllers/implement"
	"be-assignment/prisma/db"
	repositoryImplement "be-assignment/repositories/implement"
	serviceImplement "be-assignment/services/implement"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserRoutes(g *gin.Engine, gr *gin.RouterGroup, db *db.PrismaClient) {
	userRepository := repositoryImplement.RepositoryUser(db)
	userService := serviceImplement.ServiceUser(userRepository)
	userController := controllerImplement.ControllerUser(userService)

	gr.GET("/test", func(ctx *gin.Context) {
		users, err := userController.GetUsers(ctx)
		if err != nil {
			print("HEHEHEHEH")
			print(err)
		}

		ctx.JSON(http.StatusOK, gin.H{
			"message": "Hellow Transaction Test",
			"data":    users,
		})
	})
}
