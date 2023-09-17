package schema

import "github.com/graphql-go/graphql"

var (
	// AuthorList []Author
	BookList []Book
)

type Book struct {
	ID        string   `json:"id"`
	Title     string   `json:"title"`
	Authors   []string `json:"book_authors"`
	Author    []Author `json:"book_author_type"`
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
		"book_authors": &graphql.Field{
			Type: graphql.NewList(graphql.String),
		},
		"book_author_type": &graphql.Field{
			Type: graphql.NewList(AuthorType),
		},
	},
})
