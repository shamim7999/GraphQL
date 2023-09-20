package queries

import (
	"graphql_test/resolvers"
	"graphql_test/schema"

	"github.com/graphql-go/graphql"
)

var RootQuery = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootQuery",
	Fields: graphql.Fields{
		"GetBooksByAuthorName": &graphql.Field{
			Type:        schema.AuthorType,
			Description: "List of books with associated todos",
			Args: graphql.FieldConfigArgument{
				"author_name": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: resolvers.GetBooksByAuthorName,
		},

		"GetBooks": &graphql.Field{
			Type:        graphql.NewList(schema.BookType),
			Description: "Returns  all books",
			Args:        graphql.FieldConfigArgument{},
			Resolve:     resolvers.GetBooks,
		},

		"GetAuthors": &graphql.Field{
			Type:        graphql.NewList(schema.AuthorType),
			Description: "Returns  all Authors",
			Args:        graphql.FieldConfigArgument{},
			Resolve:     resolvers.GetAuthors,
		},

		"GetAuthorByName": &graphql.Field{
			Type:        schema.AuthorType,
			Description: "Returns Authors By Name",
			Args: graphql.FieldConfigArgument{
				"author_name": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: resolvers.GetAuthorByName,
		},
	},
})
