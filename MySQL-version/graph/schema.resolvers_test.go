package graph

import (
	"context"
	"imdb/graph/model"
	"reflect"
	"testing"
)

func Test_mutationResolver_CreateMovie(t *testing.T) {

	tests := []struct {
		name     string
		Resolver *Resolver
		input    model.NewMovie
		want     *model.Movie
		wantErr  bool
	}{
		{
			name:     "test",
			Resolver: mutationResolver,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &mutationResolver{
				Resolver: tt.Resolver,
			}
			got, err := r.CreateMovie(context.Background(), tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("mutationResolver.CreateMovie() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mutationResolver.CreateMovie() = %v, want %v", got, tt.want)
			}
		})
	}
}
