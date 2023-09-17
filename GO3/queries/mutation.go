package queries

import (
	"graphql_test/resolver"
	"graphql_test/schema"

	"github.com/graphql-go/graphql"
)

var RootMutation = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootMutation",
	Fields: graphql.Fields{
		"CreateNewAuthor": &graphql.Field{
			Type:        schema.AuthorType,
			Description: "Create new Author",
			Args:        graphql.FieldConfigArgument{},
			Resolve:     resolver.CreateNewAuthor,
		},
	},
})
