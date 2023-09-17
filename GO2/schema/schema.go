package schema

import (
	"github.com/graphql-go/graphql"
	"graphql_test/helper"
	"graphql_test/resolver"
)

var rootQuery = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootQuery",
	Fields: graphql.Fields{
		///////////////////////////
		"Getbooks_by_authid": &graphql.Field{
			Type:        graphql.NewList(helper.BookType),
			Description: "List of books with associated todos",
			Args: graphql.FieldConfigArgument{
				"author_id": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: resolver.Getbooks_by_authid,
		},

		"Getbooks": &graphql.Field{
			Type:        graphql.NewList(helper.BookType),
			Description: "Returns  all books",
			Args:        graphql.FieldConfigArgument{},
			Resolve:     resolver.Getbooks,
		},

		//////////////////////////

	},
})

var rootMutation = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootMutation",
	Fields: graphql.Fields{
		"createTodo": &graphql.Field{
			Type:        helper.TodoType,
			Description: "Create new todo",
			Args:        graphql.FieldConfigArgument{},
			Resolve:     resolver.Create_todo,
		},
	},
})
var TodoSchema, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query:    rootQuery,
	Mutation: rootMutation,
})
