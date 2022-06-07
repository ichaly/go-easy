package core

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

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
	panic(fmt.Errorf("not implemented"))
}

// Team returns TeamResolver implementation.
func (r *Resolver) Team() TeamResolver { return &teamResolver{r} }

type teamResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *teamResolver) User(ctx context.Context, obj *Team) (*User, error) {
	panic(fmt.Errorf("not implemented"))
}
