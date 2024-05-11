package main

import (
	config "be-assignment/configs"
	"be-assignment/routes"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/supertokens/supertokens-golang/supertokens"
)

func main() {
	g := gin.Default()

	db := config.ConnectDB()

	routes.RouteInit(g, db)

	g.Use(func(ctx *gin.Context) {
		config.InitSupertokens(ctx, db)

		supertokens.Middleware(http.HandlerFunc(
			func(rw http.ResponseWriter, r *http.Request) {
				ctx.Next()
			})).ServeHTTP(ctx.Writer, ctx.Request)
		// we call Abort so that the next handler in the chain is not called, unless we call Next explicitly
		ctx.Abort()
	})

	// Listen and Server in 0.0.0.0:8080
	g.Run(":8080")
}
