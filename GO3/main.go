package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/graphql-go/graphql"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"graphql_test/db"
	"graphql_test/queries"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"time"
)

//func init() {
//
//
//}

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

	db.ClientOptions = options.Client().ApplyURI(
		"mongodb://admin:secret@localhost:27017",
	).SetDirect(true)

	db.CtxBook = context.Background()
	db.ClientBook, db.ErrBook = mongo.Connect(db.CtxBook, db.ClientOptions)
	if db.ErrBook != nil {
		log.Fatal(db.ErrBook)
	}
	defer db.ClientBook.Disconnect(db.CtxBook)

	db.CtxAuthor = context.Background()
	db.ClientAuthor, db.ErrBook = mongo.Connect(db.CtxAuthor, db.ClientOptions)
	if db.ErrAuthor != nil {
		log.Fatal(db.ErrAuthor)
	}
	defer db.ClientAuthor.Disconnect(db.CtxAuthor)

	fmt.Println("Connected to MongoDB")

	db.Database = db.ClientBook.Database("shamim")
	db.CollectionBook = db.Database.Collection("Book")
	db.CollectionAuthor = db.Database.Collection("Author")
	rand.Seed(time.Now().UnixNano())

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
