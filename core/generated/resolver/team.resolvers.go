package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/ichaly/go-easy/core"
	"github.com/ichaly/go-easy/core/generated"
	"github.com/ichaly/go-easy/core/generated/model"
)

func (r *mutationResolver) CreateTeam(ctx context.Context, input model.NewTeam) (*core.Team, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Teams(ctx context.Context) ([]*core.Team, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *teamResolver) ID(ctx context.Context, obj *core.Team) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *teamResolver) User(ctx context.Context, obj *core.Team) (*core.User, error) {
	panic(fmt.Errorf("not implemented"))
}

// Team returns generated.TeamResolver implementation.
func (r *Resolver) Team() generated.TeamResolver { return &teamResolver{r} }

type teamResolver struct{ *Resolver }
