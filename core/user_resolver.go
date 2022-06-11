package core

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"strconv"

	"github.com/ichaly/go-easy/core/model"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*User, error) {
	return r.userService.SignIn(ctx, input.Name, input.Password)
}

func (r *queryResolver) Users(ctx context.Context) ([]User, error) {
	return r.userService.ListAll(ctx)
}

func (r *userResolver) ID(ctx context.Context, obj *User) (string, error) {
	return strconv.FormatUint(obj.ID, 10), nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

// User returns UserResolver implementation.
func (r *Resolver) User() UserResolver { return &userResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type userResolver struct{ *Resolver }
