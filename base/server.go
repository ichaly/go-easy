package base

import (
	"context"
	"fmt"
	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/ichaly/go-env"
	"go.uber.org/fx"
)

func playgroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL", "/api")
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

func NewServer(lifecycle fx.Lifecycle, schema graphql.ExecutableSchema) {
	r := gin.New()
	r.Use(cors.Default())
	r.GET("/", playgroundHandler())
	r.POST("/api", graphqlHandler(schema))
	lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			port, _ := env.String(":${PORT:=8080}")
			fmt.Printf("Now server is running on port %s\n", port)
			fmt.Printf("Connect to http://localhost%s/ for GraphQL playground\n", port)
			fmt.Printf("Test with Get: curl -g 'http://localhost%s/api?query={hello}'\n", port)
			_ = r.Run(port)
			return nil
		},
	})
}
