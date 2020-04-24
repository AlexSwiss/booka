package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/AlexSwiss/bookworm/graph/generated"
	"github.com/AlexSwiss/bookworm/graph/model"
	"github.com/AlexSwiss/bookworm/graph/models"
)

func (r *bookResolver) Category(ctx context.Context, obj *models.Book) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) AddBook(ctx context.Context, input *model.NewBook, author []*model.NewAuthor) (*models.Book, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Books(ctx context.Context) ([]*models.Book, error) {
	panic(fmt.Errorf("not implemented"))
}

// Book returns generated.BookResolver implementation.
func (r *Resolver) Book() generated.BookResolver { return &bookResolver{r} }

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type bookResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
