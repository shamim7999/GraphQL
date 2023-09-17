package schema

import "github.com/graphql-go/graphql"

var (
	AuthorList []Author
	//BookList []Book
)

type Author struct {
	ID      		string `json:"id"`
	AuthorName 		string `json:"author_name"`
}

var AuthorType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Author",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.String,
		},
		"author_name": &graphql.Field{
			Type: graphql.String,
		},
	},
})
