package resolver

import (
	"fmt"
	"github.com/graphql-go/graphql"
	"graphql_test/helper"
	"math/rand"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func Getbooks_by_authid(p graphql.ResolveParams) (interface{}, error) {
	authorID, isOK := p.Args["author_id"].(string)
	if !isOK {
		return nil, nil
	}

	var matchingTodos []helper.Todo
	var boi []helper.Book

	for _, ids := range helper.TodoList {
		if ids.ID == authorID {
			matchingTodos = append(matchingTodos, ids)
		}
	}
	boi = append(boi, helper.Book{
		Todo:      matchingTodos,
		Author_Id: authorID,
	})
	return boi, nil
}

func Getbooks(p graphql.ResolveParams) (interface{}, error) {
	return helper.BookList, nil
}

func Create_todo(params graphql.ResolveParams) (interface{}, error) {

	newID := RandStringRunes(8)

	var (
		texts   string
		authors string
	)

	fmt.Printf("Enter a Author Name: ")

	_, err := fmt.Scanf("%v\n", &authors)
	if err != nil {
		fmt.Printf("Error: ", err)
		return nil, nil
	}

	fmt.Printf("Enter some Texts: ")

	_, err = fmt.Scanf("%v\n", &texts)
	if err != nil {
		fmt.Printf("Error: ", err)
		return nil, nil
	}
	newTodo := helper.Todo{
		ID:     newID,
		Text:   texts,
		Author: authors,
	}

	helper.TodoList = append(helper.TodoList, newTodo)
	return newTodo, nil
}
