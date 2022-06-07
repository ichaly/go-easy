package core

//go:generate go run github.com/99designs/gqlgen generate
import "github.com/99designs/gqlgen/graphql"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	userService IUserService
	teamService ITeamService
}

func NewSchema(
	userService IUserService,
	teamService ITeamService,
) graphql.ExecutableSchema {
	return NewExecutableSchema(Config{Resolvers: &Resolver{
		userService,
		teamService,
	}})
}
