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
		"Getbooks_by_authid": &graphql.Field{
			Type:        graphql.NewList(schema.BookType),
			Description: "List of books with associated todos",
			Args: graphql.FieldConfigArgument{
				"author_id": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: resolver.Getbooks_by_authid,
		},

		"Getbooks": &graphql.Field{
			Type:        graphql.NewList(schema.BookType),
			Description: "Returns  all books",
			Args:        graphql.FieldConfigArgument{},
			Resolve:     resolver.Getbooks,
		},

		//////////////////////////

	},
})
