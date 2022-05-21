package main

import (
	"fmt"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/a8m/envsubst"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/ichaly/go-easy/core/graph"
	"github.com/ichaly/go-easy/core/graph/generated"
)

func main() {
	port, _ := envsubst.String(":${PORT:=8080}")
	r := gin.New()
	r.Use(cors.Default())
	r.POST("/api", func() gin.HandlerFunc {
		h := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))
		return func(c *gin.Context) {
			h.ServeHTTP(c.Writer, c.Request)
		}
	}())
	fmt.Println("Now server is running on port 8080")
	fmt.Println("Test with Get      : curl -g 'http://localhost:8080/api?query={hello}'")
	_ = r.Run(port)
}
