package main

import (
	"encoding/json"
	"fmt"
	"graphql_test/queries"
	"graphql_test/schema"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"

	"github.com/graphql-go/graphql"
)

func init() {

	book1 := &schema.Book{ID: "a", Title: "Golang", Authors: []string{"Shamim", "Saiful", "Ashraful", "Neaj"}}
	book2 := &schema.Book{ID: "b2", Title: "C/C++", Authors: []string{"Ashraful", "Neaj"}}
	book3 := &schema.Book{ID: "c3", Title: "JAVA", Authors: []string{"Ashraful", "Neaj"}}
	book4 := &schema.Book{ID: "d4", Title: "Python", Authors: []string{"Shamim", "Saiful"}}
	schema.BookList = append(schema.BookList, book1, book2, book3, book4)

	Authors1 := schema.Author{ID: "a", AuthorName: "Shamim", Book: schema.BookList}
	Authors2 := schema.Author{ID: "b", AuthorName: "Saiful", Book: schema.BookList}
	Authors3 := schema.Author{ID: "c", AuthorName: "Ashraful", Book: schema.BookList}

	schema.AuthorList = append(schema.AuthorList, Authors1, Authors2, Authors3)

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
		result := executeQuery(t.Query, queries.QueriesAndMutation)
		json.NewEncoder(w).Encode(result)
	})
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)
	fmt.Println("Now server is running on port 8080")
	fmt.Println("Get single Authors: curl -g 'http://localhost:8080/graphql?query={Authors(id:\"b\"){id,text,done}}'")
	fmt.Println("Create new Authors: curl -g 'http://localhost:8080/graphql?query=mutation+_{createAuthors(text:\"My+new+Authors\"){id,text,done}}'")
	fmt.Println("Update Authors: curl -g 'http://localhost:8080/graphql?query=mutation+_{updateAuthors(id:\"a\",done:true){id,text,done}}'")
	fmt.Println("Load Authors list: curl -g 'http://localhost:8080/graphql?query={AuthorList{id,text,done}}'")
	fmt.Println("Access the web app via browser at 'http://localhost:8080'")

	http.ListenAndServe(":8080", nil)
}
