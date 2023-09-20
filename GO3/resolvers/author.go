package resolvers

import (
	"fmt"
	"github.com/graphql-go/graphql"
	"go.mongodb.org/mongo-driver/bson"
	db "graphql_test/db/queries"
	"graphql_test/schema"
	"log"
)

func GetAuthors(p graphql.ResolveParams) (interface{}, error) {
	filter := bson.M{}

	var AuthorsList []schema.Author

	AuthorsList = db.GetDataFromAuthorCollection(filter)

	if len(AuthorsList) == 0 {
		return nil, nil
	}
	return AuthorsList, nil
}

func CreateNewAuthor(p graphql.ResolveParams) (interface{}, error) {
	filter := bson.M{}
	authorName, isOK := p.Args["author_name"].(string)
	if !isOK {
		return nil, nil
	}
	authorID := schema.RandStringRunes(8)

	var BooksList []*schema.Book
	BooksList = db.GetDataFromBookCollection(filter)

	newAuthor := schema.Author{
		ID:         authorID,
		AuthorName: authorName,
		Book:       BooksList,
	}

	err := db.InsertAuthor(newAuthor)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Author inserted successfully.")
	return newAuthor, nil
}

func GetAuthorByName(p graphql.ResolveParams) (interface{}, error) {
	authorName, isOK := p.Args["author_name"].(string)
	if !isOK {
		return nil, nil
	}

	filter := bson.M{
		"authorname": authorName,
	}

	var AuthorsList []schema.Author
	AuthorsList = db.GetDataFromAuthorCollection(filter)

	if len(AuthorsList) == 0 {
		return nil, nil
	}
	return AuthorsList[0], nil
}
