package main

import (
	"encoding/json"
	"fmt"
	"graphql_test/helper"
	"graphql_test/schema"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"

	"github.com/graphql-go/graphql"
)

func init() {
	todo1 := helper.Todo{ID: "a", Text: "A todo not to forget", Author: "Shamim"}
	todo2 := helper.Todo{ID: "b", Text: "This is the most important", Author: "Saiful"}
	todo3 := helper.Todo{ID: "b", Text: "Please do this or else", Author: "Saiful"}
	todo4 := helper.Todo{ID: "c", Text: "Please do this or else", Author: "Ashraful"}
	todo5 := helper.Todo{ID: "a", Text: "Please do this or else", Author: "Shamim"}
	todo6 := helper.Todo{ID: "a", Text: "Please do this or else", Author: "Shamim"}
	helper.TodoList = append(helper.TodoList, todo1, todo2, todo3, todo4, todo5, todo6)

	book1 := helper.Book{ID: "a1", Title: "A todo not to forget", Author_Id: "a", Todo: helper.TodoList}
	book2 := helper.Book{ID: "b2", Title: "This is the most important", Author_Id: "a", Todo: helper.TodoList}
	book3 := helper.Book{ID: "c3", Title: "Please do this or else", Author_Id: "b", Todo: helper.TodoList}
	book4 := helper.Book{ID: "d4", Title: "Please do this or else", Author_Id: "c", Todo: helper.TodoList}
	helper.BookList = append(helper.BookList, book1, book2, book3, book4)
	rand.Seed(time.Now().UnixNano())
}

func executeQuery(query string, schema graphql.Schema) *graphql.Result {
	result := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: query,
	})
	if len(result.Errors) > 0 {
		fmt.Printf("wrong result, unexpected errors: %v", result.Errors)
	}
	return result
}

type ReqBody struct {
	Query string `json:"query"`
}

func main() {
	http.HandleFunc("/graphql", func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			panic(err)
		}
		var t ReqBody
		err = json.Unmarshal(body, &t)
		if err != nil {
			panic(err)
		}
		result := executeQuery(t.Query, schema.TodoSchema)
		json.NewEncoder(w).Encode(result)
	})
	// Serve static files
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)
	// Display some basic instructions
	fmt.Println("Now server is running on port 8080")
	fmt.Println("Get single todo: curl -g 'http://localhost:8080/graphql?query={todo(id:\"b\"){id,text,done}}'")
	fmt.Println("Create new todo: curl -g 'http://localhost:8080/graphql?query=mutation+_{createTodo(text:\"My+new+todo\"){id,text,done}}'")
	fmt.Println("Update todo: curl -g 'http://localhost:8080/graphql?query=mutation+_{updateTodo(id:\"a\",done:true){id,text,done}}'")
	fmt.Println("Load todo list: curl -g 'http://localhost:8080/graphql?query={todoList{id,text,done}}'")
	fmt.Println("Access the web app via browser at 'http://localhost:8080'")

	http.ListenAndServe(":8080", nil)
}
