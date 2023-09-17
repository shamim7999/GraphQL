package queries

import (
	"github.com/graphql-go/graphql"
)

var Todo, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query:    RootQuery,
	Mutation: RootMutation,
})
