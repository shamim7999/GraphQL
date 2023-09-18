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

func CreateNewBook(p graphql.ResolveParams) (interface{}, error) {
	newBook := schema.Book{
		ID:      RandStringRunes(8),
		Title:   "A new Book",
		Authors: []string{"Ashraful", "Ashik", "Morshed"},
		Author:  schema.AuthorList,
	}
	schema.BookList = append(schema.BookList, newBook)
	return newBook, nil
}

func CreateNewBookByParameter(p graphql.ResolveParams) (interface{}, error) {
	newID := RandStringRunes(8)
	newTitle, isOK := p.Args["title"].(string)
	if !isOK {
		return nil, nil
	}
	bookAuthors, isOK := p.Args["book_authors"].([]interface{})
	if !isOK {
		return nil, nil
	}

	// Convert the list of interfaces to a list of strings
	var authors []string
	for _, author := range bookAuthors {
		if authorStr, isStr := author.(string); isStr {
			authors = append(authors, authorStr)
		}
	}

	newBook := schema.Book{
		ID:      newID,
		Title:   newTitle,
		Authors: authors,
		Author:  schema.AuthorList,
	}
	schema.BookList = append(schema.BookList, newBook)
	return newBook, nil
}
