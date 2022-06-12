package base

import (
	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

const ENDPOINT string = "/api"

func playgroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL", ENDPOINT)
	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func graphqlHandler(schema graphql.ExecutableSchema) gin.HandlerFunc {
	h := handler.NewDefaultServer(schema)
	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func NewEngine(schema graphql.ExecutableSchema) *gin.Engine {
	r := gin.New()
	r.Use(cors.Default())
	r.POST(ENDPOINT, graphqlHandler(schema))
	r.GET("/", playgroundHandler())
	return r
}
