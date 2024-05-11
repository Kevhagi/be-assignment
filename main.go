package main

import (
	config "be-assignment/configs"
	"be-assignment/routes"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/supertokens/supertokens-golang/supertokens"
)

func main() {
	g := gin.Default()

	db := config.ConnectDB()

	g.Use(func(ctx *gin.Context) {
		config.InitSupertokens(ctx, db)

		supertokens.Middleware(http.HandlerFunc(
			func(rw http.ResponseWriter, r *http.Request) {
				ctx.Next()
			})).ServeHTTP(ctx.Writer, ctx.Request)
		// we call Abort so that the next handler in the chain is not called, unless we call Next explicitly
		ctx.Abort()
	})

	routes.RouteInit(g, db)

	if err := godotenv.Load(); err != nil {
		panic("Failed to load env file")
	}

	if PORT := os.Getenv("PORT"); PORT != "" {
		// Handling specified PORT in .env file
		err := g.Run(":" + PORT)
		if err != nil {
			panic("[Error] failed to start Gin server due to: " + err.Error())
		}
	} else {
		// Handling default PORT by Gin
		err := g.Run()
		if err != nil {
			panic("[Error] failed to start Gin server due to: " + err.Error())
		}
	}
}
