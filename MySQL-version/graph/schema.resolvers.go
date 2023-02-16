package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.24

import (
	"context"
	"fmt"
	"imdb/graph/model"
)

// CreateMovie is the resolver for the createMovie field.
func (r *mutationResolver) CreateMovie(ctx context.Context, input model.NewMovie) (*model.Movie, error) {
	panic(fmt.Errorf("not implemented: CreateMovie - createMovie"))
}

// CreateActor is the resolver for the createActor field.
func (r *mutationResolver) CreateActor(ctx context.Context, input model.NewActor) (*model.Actor, error) {
	panic(fmt.Errorf("not implemented: CreateActor - createActor"))
}

// Actors is the resolver for the actors field.
func (r *queryResolver) Actors(ctx context.Context) ([]*model.Actor, error) {
	panic(fmt.Errorf("not implemented: Actors - actors"))
}

// Movies is the resolver for the movies field.
func (r *queryResolver) Movies(ctx context.Context) ([]*model.Movie, error) {
	panic(fmt.Errorf("not implemented: Movies - movies"))
}

// Actor is the resolver for the actor field.
func (r *queryResolver) Actor(ctx context.Context, input *model.FetchActor) (*model.Actor, error) {
	panic(fmt.Errorf("not implemented: Actor - actor"))
}

// Movie is the resolver for the movie field.
func (r *queryResolver) Movie(ctx context.Context, input *model.FetchMovie) (*model.Movie, error) {
	panic(fmt.Errorf("not implemented: Movie - movie"))
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
