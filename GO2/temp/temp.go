package schema

import (
	"fmt"
	"math/rand"

	"github.com/graphql-go/graphql"
)

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

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

// define custom GraphQL ObjectType `todoType` for our Golang struct `Todo`
// Note that
// - the fields in our todoType maps with the json tags for the fields in our struct
// - the field type matches the field type in our struct
var todoType = graphql.NewObject(graphql.ObjectConfig{
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

var bookType = graphql.NewObject(graphql.ObjectConfig{
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
			Type: graphql.NewList(todoType),
		},
	},
})

// root mutation
var rootMutation = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootMutation",
	Fields: graphql.Fields{
		/*
		   curl -g 'http://localhost:8080/graphql?query=mutation+_{createTodo(text:"My+new+todo"){id,text,done}}'
		*/
		"createTodo": &graphql.Field{
			Type:        todoType, // the return type for this field
			Description: "Create new todo",
			Args: graphql.FieldConfigArgument{
				"text": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {

				// marshall and cast the argument value
				text, _ := params.Args["text"].(string)

				// figure out new id
				newID := RandStringRunes(8)

				// perform mutation operation here
				// for e.g. create a Todo and save to DB.
				newTodo := Todo{
					ID:     newID,
					Text:   text,
					Author: "Shamim",
				}

				TodoList = append(TodoList, newTodo)

				// return the new Todo object that we supposedly save to DB
				// Note here that
				// - we are returning a `Todo` struct instance here
				// - we previously specified the return Type to be `todoType`
				// - `Todo` struct maps to `todoType`, as defined in `todoType` ObjectConfig`
				return newTodo, nil
			},
		},
		/*
		   curl -g 'http://localhost:8080/graphql?query=mutation+_{updateTodo(id:"a",done:true){id,text,done}}'
		*/
		"updateTodo": &graphql.Field{
			Type:        todoType, // the return type for this field
			Description: "Update existing todo, mark it done or not done",
			Args: graphql.FieldConfigArgument{
				"author": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				// marshall and cast the argument value
				author, _ := params.Args["author"].(string)
				id, _ := params.Args["id"].(string)
				affectedTodo := Todo{}

				// Search list for todo with id and change the done variable
				for i := 0; i < len(TodoList); i++ {
					if TodoList[i].ID == id {
						TodoList[i].Author = author
						// Assign updated todo so we can return it
						affectedTodo = TodoList[i]
						break
					}
				}
				// Return affected todo
				return affectedTodo, nil
			},
		},
	},
})

// root query
// we just define a trivial example here, since root query is required.
// Test with curl
// curl -g 'http://localhost:8080/graphql?query={lastTodo{id,text,done}}'
var rootQuery = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootQuery",
	Fields: graphql.Fields{

		/*
		   curl -g 'http://localhost:8080/graphql?query={todo(id:"b"){id,text,done}}'
		*/
		"todo": &graphql.Field{
			Type:        todoType,
			Description: "Get single todo",
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {

				idQuery, isOK := params.Args["id"].(string)
				if isOK {
					// Search for el with id
					for _, todo := range TodoList {
						if todo.ID == idQuery {
							return todo, nil
						}
					}
				}

				return Todo{}, nil
			},
		},

		"lastTodo": &graphql.Field{
			Type:        todoType,
			Description: "Last todo added",
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				return TodoList[len(TodoList)-1], nil
			},
		},

		/*
		   curl -g 'http://localhost:8080/graphql?query={todoList{id,text,done}}'
		*/
		"todoList": &graphql.Field{
			Type:        graphql.NewList(todoType),
			Description: "List of todos",

			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				fmt.Println("hello")
				return TodoList, nil
			},
		},

		"getBookAndAuthorByAuthor": &graphql.Field{
			Type:        graphql.NewList(todoType),
			Description: "Shows One Author and Book",
			Args: graphql.FieldConfigArgument{

				"author": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				nameAuthor, isOK := p.Args["author"].(string)
				var matchingTodos []Todo
				if isOK {
					// Search for el with id
					for _, todo := range TodoList {
						if todo.Author == nameAuthor {
							matchingTodos = append(matchingTodos, todo)
						}
					}
				}

				return matchingTodos, nil
			},
		},

		"getBooks": &graphql.Field{
			Type:        graphql.NewList(bookType),
			Description: "List of todos",

			Resolve: func(p graphql.ResolveParams) (interface{}, error) {

				return BookList, nil
			},
		},

		///////////////////////////
		"getBooks2": &graphql.Field{
			Type:        graphql.NewList(bookType),
			Description: "List of books with associated todos",
			Args: graphql.FieldConfigArgument{
				"author_id": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				authorID, isOK := p.Args["author_id"].(string)
				if !isOK {
					return nil, nil
				}

				var matchingTodos []Todo
				var boi []Book

				for _, ids := range TodoList {
					if ids.ID == authorID {
						matchingTodos = append(matchingTodos, ids)
					}
				}
				boi = append(boi, Book{
					Todo:      matchingTodos,
					Author_Id: authorID,
				})
				return boi, nil
			},
		},

		//////////////////////////

	},
})

// define schema, with our rootQuery and rootMutation
var TodoSchema, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query:    rootQuery,
	Mutation: rootMutation,
})

//"getBookByAuthorId": &graphql.Field{
//Type:        graphql.NewList(todoType),
//Description: "Shows One Author and Book",
//Args: graphql.FieldConfigArgument{
//
//"author_id": &graphql.ArgumentConfig{
//Type: graphql.String,
//},
////"id": &graphql.ArgumentConfig{
////	Type: graphql.NewNonNull(graphql.String),
////},
//},
//Resolve: func(p graphql.ResolveParams) (interface{}, error) {
//author_id, isOK := p.Args["author_id"].(string)
//var matchingTodos []Book
////seen := make(map[Book]bool)
//if isOK {
//// Search for el with id
//for _, todo := range TodoList {
//for _, book := range BookList {
//if todo.ID == author_id {
//matchingTodos = append(matchingTodos, book)
////seen[book] = true
//}
//}
//}
//}
//
//return matchingTodos, nil
//},
//},

//func getBookAndAuthorByAuthor2(p graphql.ResolveParams) (interface{}, error) {
//	nameAuthor, isOK := p.Args["author"].(string)
//	var matchingTodos []Todo
//	if isOK {
//		// Search for el with id
//		for _, todo := range TodoList {
//			if todo.Author == nameAuthor {
//				matchingTodos = append(matchingTodos, todo)
//			}
//		}
//	}
//
//	return matchingTodos, nil
//}