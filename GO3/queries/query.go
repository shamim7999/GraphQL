package queries

import (
	"graphql_test/resolver"
	"graphql_test/schema"

	"github.com/graphql-go/graphql"
)

var RootQuery = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootQuery",
	Fields: graphql.Fields{
		///////////////////////////
		"GetBooksByAuthorName": &graphql.Field{
			Type:        graphql.NewList(schema.BookType),
			Description: "List of books with associated todos",
			Args: graphql.FieldConfigArgument{
				"author_name": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: resolver.GetBooksByAuthorName,
		},

		"GetBooks": &graphql.Field{
			Type:        graphql.NewList(schema.BookType),
			Description: "Returns  all books",
			Args:        graphql.FieldConfigArgument{},
			Resolve:     resolver.GetBooks,
		},

		"GetAuthors": &graphql.Field{
			Type:        graphql.NewList(schema.AuthorType),
			Description: "Returns  all Authors",
			Args:        graphql.FieldConfigArgument{},
			Resolve:     resolver.GetAuthors,
		},

		"GetAuthorByName": &graphql.Field{
			Type:        schema.AuthorType,
			Description: "Returns  all Authors",
			Args: graphql.FieldConfigArgument{
				"author_name": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: resolver.GetAuthorByName,
		},

		//////////////////////////

	},
})
