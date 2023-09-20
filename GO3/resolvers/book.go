package resolvers

import (
	"fmt"
	"github.com/graphql-go/graphql"
	"go.mongodb.org/mongo-driver/bson"
	db "graphql_test/db/queries"
	"graphql_test/schema"
	"log"
)

func GetBooks(p graphql.ResolveParams) (interface{}, error) {
	filter := bson.M{}

	var BooksList []*schema.Book

	BooksList = db.GetDataFromBookCollection(filter)

	return BooksList, nil
}

func GetBooksByAuthorName(p graphql.ResolveParams) (interface{}, error) {

	authorName, isOK := p.Args["author_name"].(string)
	if !isOK {
		return nil, nil
	}

	var AuthorsList []schema.Author
	var BooksList []*schema.Book
	BooksByAuthorName := schema.Author{}

	filter := bson.M{
		"authorname": authorName,
	}
	//ok = true
	AuthorsList = db.GetDataFromAuthorCollection(filter)

	filter = bson.M{
		"authors": bson.M{
			"$elemMatch": bson.M{
				"$eq": authorName,
			},
		},
	}

	//ok = false
	BooksList = db.GetDataFromBookCollection(filter)

	if len(AuthorsList) == 0 {
		return nil, nil
	}

	BooksByAuthorName = schema.Author{
		ID:         AuthorsList[0].ID,
		AuthorName: AuthorsList[0].AuthorName,
		Book:       BooksList,
	}

	return BooksByAuthorName, nil
}

func CreateNewBook(p graphql.ResolveParams) (interface{}, error) {
	newID := schema.RandStringRunes(8)
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

	fmt.Println("Book inserted successfully.")
	err := db.InsertBook(newBook)
	if err != nil {
		log.Fatal(err)
	}
	return newBook, nil
}
