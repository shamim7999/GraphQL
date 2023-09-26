package schema

import (
	"context"
	"fmt"
	"github.com/graph-gophers/dataloader"
	"github.com/graphql-go/graphql"
	"go.mongodb.org/mongo-driver/bson"
	"graphql_test/db"
	"math/rand"
)

type Book struct {
	ID      string   `json:"id" bson:"id"`
	Title   string   `json:"title" bson:"title"`
	Authors []string `json:"authors" bson:"authors"`
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

var BookType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Book",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.String,
		},
		"title": &graphql.Field{
			Type: graphql.String,
		},
		"authors": &graphql.Field{
			Type: graphql.NewList(graphql.String),
		},
	},
})

var BookLoader = dataloader.NewBatchedLoader(func(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
	var bookIDs []string
	for _, key := range keys {
		bookID := key.String()
		bookIDs = append(bookIDs, bookID)
	}

	var books []*Book // Use the Book type from your schema package
	filter := bson.M{"_id": bson.M{"$in": bookIDs}}
	cursor, err := db.CollectionBook.Find(ctx, filter)
	if err != nil {
		results := make([]*dataloader.Result, len(keys))
		for i := range results {
			results[i] = &dataloader.Result{Error: err}
		}
		return results
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var book Book
		if err := cursor.Decode(&book); err != nil {
			results := make([]*dataloader.Result, len(keys))
			for i := range results {
				results[i] = &dataloader.Result{Error: err}
			}
			return results
		}
		books = append(books, &book)
	}

	results := make([]*dataloader.Result, len(keys))
	for i, key := range keys {
		bookID := key.String()
		var matchingBook *Book
		for _, book := range books {
			if book.ID == bookID {
				matchingBook = book
				break
			}
		}
		if matchingBook != nil {
			results[i] = &dataloader.Result{Data: matchingBook}
		} else {
			results[i] = &dataloader.Result{Error: fmt.Errorf("Book not found: %s", bookID)}
		}
	}

	return results
})
