package schema

import "github.com/graphql-go/graphql"

var (
	AuthorList []Author
	//BookList []Book
)

type Author struct {
	ID      string `json:"id"`
	Text    string `json:"text"`
	Authors string `json:"authors"`
}

var AuthorType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Author",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.String,
		},
		"text": &graphql.Field{
			Type: graphql.String,
		},
		"authors": &graphql.Field{
			Type: graphql.String,
		},
	},
})
