package main

import (
	config "be-assignment/config"
	"be-assignment/routes"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/supertokens/supertokens-golang/supertokens"
)

func main() {
	g := gin.Default()

	routes.RouteInit(g.Group("/"))

	config.ConnectDB()
	config.InitSupertokens()

	g.Use(func(c *gin.Context) {
		supertokens.Middleware(http.HandlerFunc(
			func(rw http.ResponseWriter, r *http.Request) {
				c.Next()
			})).ServeHTTP(c.Writer, c.Request)
		// we call Abort so that the next handler in the chain is not called, unless we call Next explicitly
		c.Abort()
	})

	// Listen and Server in 0.0.0.0:8080
	g.Run(":8080")
}
