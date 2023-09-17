package schema

import "github.com/graphql-go/graphql"

var (
	//TodoList []Todo
	BookList []Book
)

type Book struct {
	ID        string `json:"id"`
	Title     string `json:"title"`
	Author_Id string `json:"author_id"`
	Author      []Author `json:"author"`
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
		"todo": &graphql.Field{
			Type: graphql.NewList(AuthorType),
		},
	},
})
