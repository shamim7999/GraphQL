package schema

import (
	"github.com/graphql-go/graphql"
)

var (
	AuthorList []Author
)

type Author struct {
	ID         string  `json:"id"`
	AuthorName string  `json:"author_name"`
	Book       []*Book `json:"book"`
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
		"book": &graphql.Field{
			Type: graphql.NewList(BookType),
		},
	},
})

func GetAuthors(p graphql.ResolveParams) (interface{}, error) {
	return AuthorList, nil
}

func GetAuthorByName(p graphql.ResolveParams) (interface{}, error) {
	authorName, isOK := p.Args["author_name"].(string)
	if !isOK {
		return nil, nil
	}
	var author Author
	for _, person := range AuthorList {
		if person.AuthorName == authorName {
			author.ID = person.ID
			author.AuthorName = person.AuthorName
			author.Book = BookList
			break
		}
	}
	return author, nil
}

func CreateNewAuthor(p graphql.ResolveParams) (interface{}, error) {
	authorName, isOK := p.Args["author_name"].(string)
	if !isOK {
		return nil, nil
	}
	authorID := RandStringRunes(8)

	newAuthor := Author{
		ID:         authorID,
		AuthorName: authorName,
		Book:       BookList,
	}

	AuthorList = append(AuthorList, newAuthor)
	return newAuthor, nil
}
