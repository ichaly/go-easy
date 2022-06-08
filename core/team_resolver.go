package core

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"strconv"

	"github.com/ichaly/go-easy/core/model"
)

func (r *mutationResolver) CreateTeam(ctx context.Context, input model.NewTeam) (*Team, error) {
	t := &Team{Name: input.Name}
	_, err := r.teamService.Save(ctx, t)
	return t, err
}

func (r *queryResolver) Teams(ctx context.Context) ([]Team, error) {
	return r.teamService.ListAll(ctx)
}

func (r *teamResolver) ID(ctx context.Context, obj *Team) (string, error) {
	return strconv.FormatUint(obj.ID, 10), nil
}

// Team returns TeamResolver implementation.
func (r *Resolver) Team() TeamResolver { return &teamResolver{r} }

type teamResolver struct{ *Resolver }
