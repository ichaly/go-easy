package core

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/ichaly/go-easy/core/model"
)

func (r *mutationResolver) CreateTeam(ctx context.Context, input model.NewTeam) (*Team, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Teams(ctx context.Context) ([]*Team, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *teamResolver) ID(ctx context.Context, obj *Team) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *teamResolver) User(ctx context.Context, obj *Team) (*User, error) {
	panic(fmt.Errorf("not implemented"))
}

// Team returns TeamResolver implementation.
func (r *Resolver) Team() TeamResolver { return &teamResolver{r} }

type teamResolver struct{ *Resolver }
