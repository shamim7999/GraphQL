package queries

import (
	"github.com/graphql-go/graphql"
)

var QueriesAndMutation, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query:    RootQuery,
	Mutation: RootMutation,
})
