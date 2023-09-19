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

	var matchingBooks []*schema.Book
	BooksByAuthorName := schema.Author{}

	for _, datas := range schema.AuthorList {
		if datas.AuthorName != authorName {
			continue
		}
		BookLists := datas.Book

		for _, Books := range BookLists {
			for _, names := range Books.Authors {
				if names == authorName {
					matchingBooks = append(matchingBooks, Books)
					break
				}
			}
		}

		BooksByAuthorName = schema.Author{
			ID:         datas.ID,
			AuthorName: datas.AuthorName,
			Book:       matchingBooks,
		}

		break
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
			author.Book = schema.BookList
			break
		}
	}
	return author, nil
}

////////////////////////////////////////////////////////////////////////////////////////////

func CreateNewAuthor(p graphql.ResolveParams) (interface{}, error) {
	authorName, isOK := p.Args["author_name"].(string)
	if !isOK {
		return nil, nil
	}
	authorID := RandStringRunes(8)

	newAuthor := schema.Author{
		ID:         authorID,
		AuthorName: authorName,
		Book:       schema.BookList,
	}

	schema.AuthorList = append(schema.AuthorList, newAuthor)
	return newAuthor, nil
}

func CreateNewBook(p graphql.ResolveParams) (interface{}, error) {
	newID := RandStringRunes(8)
	newTitle, isOK := p.Args["title"].(string)
	if !isOK {
		return nil, nil
	}
	bookAuthors, isOK := p.Args["authors"].([]interface{})
	if !isOK {
		return nil, nil
	}

	var authors []string
	for _, author := range bookAuthors {
		if authorStr, isStr := author.(string); isStr {
			authors = append(authors, authorStr)
		}
	}

	newBook := &schema.Book{
		ID:      newID,
		Title:   newTitle,
		Authors: authors,
	}
	schema.BookList = append(schema.BookList, newBook)
	return newBook, nil
}
