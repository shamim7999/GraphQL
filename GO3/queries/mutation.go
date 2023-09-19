package queries

import (
	"graphql_test/schema"

	"github.com/graphql-go/graphql"
)

var RootMutation = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootMutation",
	Fields: graphql.Fields{

		"CreateNewAuthor": &graphql.Field{
			Type:        schema.AuthorType,
			Description: "Create New Author With Given Parameter",
			Args: graphql.FieldConfigArgument{
				"author_name": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"id": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: schema.CreateNewAuthor,
		},
		"CreateNewBook": &graphql.Field{
			Type:        schema.BookType,
			Description: "Create New Book by Parameter",
			Args: graphql.FieldConfigArgument{
				"title": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"authors": &graphql.ArgumentConfig{
					Type: graphql.NewList(graphql.String),
				},
			},
			Resolve: schema.CreateNewBook,
		},
	},
})
