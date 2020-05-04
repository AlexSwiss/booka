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

func (r *mutationResolver) AddBook(ctx context.Context, input *model.NewBook, authors []*model.NewAuthor) (*models.Book, error) {
	//Fetch Connection and close db
	db := models.FetchConnection()
	defer db.Close()

	//Create the book using the input structs
	book := models.Book{Name: input.Name, Category: *&input.Category}

	//initialize the author with the length of the input for author
	book.Authors = make([]*models.Author, len(authors))
	//Loop and add all items
	for index, item := range authors {
		book.Authors[index] = &models.Author{Firstname: item.Firstname, Lastname: item.Lastname}
	}
	//Create by passing the pointer to the book
	db.Create(&book)
	return &book, nil
}

func (r *mutationResolver) UpdateBook(ctx context.Context, id *int, input *model.NewBook, authors []*model.NewAuthor) (*models.Book, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteBook(ctx context.Context, id *int) ([]*models.Book, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Books(ctx context.Context, search *string) ([]*models.Book, error) {
	//Fetch a connection
	db := models.FetchConnection()
	//Defer closing the database
	defer db.Close()
	//Create an array of recipes to populate
	var books []*models.Book
	// .Preload loads the Ingredients relationship into each recipe
	db.Preload("Authors").Find(&books)
	return books, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
