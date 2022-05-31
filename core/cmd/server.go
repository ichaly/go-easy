package main

import (
	"fmt"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/ichaly/go-easy/core/graph"
	"github.com/ichaly/go-easy/core/graph/generated"
	"github.com/ichaly/go-env"
)

func playgroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL", "/api")
	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func graphqlHandler() gin.HandlerFunc {
	h := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))
	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func main() {
	port, _ := env.String(":${PORT:=8080}")
	r := gin.New()
	r.Use(cors.Default())
	r.GET("/", playgroundHandler())
	r.POST("/api", graphqlHandler())
	fmt.Println("Now server is running on port 8080")
	fmt.Printf("Connect to http://localhost:%s/ for GraphQL playground\n", port)
	fmt.Println("Test with Get: curl -g 'http://localhost:8080/api?query={hello}'")
	_ = r.Run(port)
}
