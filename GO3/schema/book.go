package schema

import "github.com/graphql-go/graphql"

var (
	//AuthorList []Author
	BookList []Book
)

type Book struct {
	ID        string   `json:"id"`
	Title     string   `json:"title"`
	Author_Id string   `json:"author_id"`
	Author    []Author `json:"book_author"`
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
		"author_id": &graphql.Field{
			Type: graphql.String,
		},
		"book_author": &graphql.Field{
			Type: graphql.NewList(AuthorType),
		},
	},
})
