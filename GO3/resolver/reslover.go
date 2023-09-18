package resolver

import (
	"graphql_test/schema"
	"math/rand"

	"github.com/graphql-go/graphql"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func GetBooksByAuthorName(p graphql.ResolveParams) (interface{}, error) {
	authorName, isOK := p.Args["author_name"].(string)
	if !isOK {
		return nil, nil
	}

	//var matchingAuthors []schema.Author
	var BooksByAuthorName []schema.Book

	for _, datas := range schema.BookList {
		for _, name := range datas.Authors {
			if name == authorName {
				BooksByAuthorName = append(BooksByAuthorName, schema.Book{
					ID:      datas.ID,
					Title:   datas.Title,
					Authors: datas.Authors,
				})
				break
			}
		}
	}

	return BooksByAuthorName, nil
}

func GetBooks(p graphql.ResolveParams) (interface{}, error) {
	return schema.BookList, nil
}

func GetAuthors(p graphql.ResolveParams) (interface{}, error) {
	return schema.AuthorList, nil
}

func GetAuthorByName(p graphql.ResolveParams) (interface{}, error) {
	authorName, isOK := p.Args["author_name"].(string)
	if !isOK {
		return nil, nil
	}
	var author schema.Author
	for _, person := range schema.AuthorList {
		if person.AuthorName == authorName {
			author.ID = person.ID
			author.AuthorName = person.AuthorName
			break
		}
	}
	return author, nil
}

////////////////////////////////////////////////////////////////////////////////////////////

func CreateNewAuthor(params graphql.ResolveParams) (interface{}, error) {

	newID := RandStringRunes(8)
	authors := "Shamim"
	newAuthor := schema.Author{
		ID:         newID,
		AuthorName: authors,
	}

	schema.AuthorList = append(schema.AuthorList, newAuthor)
	return newAuthor, nil
}

func CreateNewAuthorByParameter(p graphql.ResolveParams) (interface{}, error) {
	authorName, isOK := p.Args["author_name"].(string)
	if !isOK {
		return nil, nil
	}
	authorID, isOK := p.Args["id"].(string)
	if !isOK {
		return nil, nil
	}

	newAuthor := schema.Author{
		ID:         authorID,
		AuthorName: authorName,
	}

	schema.AuthorList = append(schema.AuthorList, newAuthor)
	return newAuthor, nil
}
