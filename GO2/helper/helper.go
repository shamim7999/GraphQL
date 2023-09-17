package helper

import "github.com/graphql-go/graphql"

var (
	TodoList []Todo
	BookList []Book
)

type Todo struct {
	ID     string `json:"id"`
	Text   string `json:"text"`
	Author string `json:"author"`
}

type Book struct {
	ID        string `json:"id"`
	Title     string `json:"title"`
	Author_Id string `json:"author_id"`
	Todo      []Todo `json:"todo"`
}

var TodoType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Todo",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.String,
		},
		"text": &graphql.Field{
			Type: graphql.String,
		},
		"author": &graphql.Field{
			Type: graphql.String,
		},
	},
})

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
			Type: graphql.NewList(TodoType),
		},
	},
})
