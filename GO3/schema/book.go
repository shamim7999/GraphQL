package schema

import (
	"github.com/graphql-go/graphql"
	"math/rand"
)

var (
	BookList []*Book
)

type Book struct {
	ID      string   `json:"id"`
	Title   string   `json:"title"`
	Authors []string `json:"authors"`
}

var BookType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Book",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.String,
		},
		"title": &graphql.Field{
			Type: graphql.String,
		},
		"authors": &graphql.Field{
			Type: graphql.NewList(graphql.String),
		},
	},
})

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

	var matchingBooks []*Book
	BooksByAuthorName := Author{}

	for _, datas := range AuthorList {
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

		BooksByAuthorName = Author{
			ID:         datas.ID,
			AuthorName: datas.AuthorName,
			Book:       matchingBooks,
		}

		break
	}

	return BooksByAuthorName, nil
}

func GetBooks(p graphql.ResolveParams) (interface{}, error) {
	return BookList, nil
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

	newBook := &Book{
		ID:      newID,
		Title:   newTitle,
		Authors: authors,
	}
	BookList = append(BookList, newBook)
	return newBook, nil
}
