package schema

import (
	"context"
	"fmt"
	"github.com/graphql-go/graphql"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"graphql_test/db"
	"log"
)

var (
	AuthorList []Author
	ok         bool
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

func GetDataFromCollection(Col *mongo.Collection, Ctx context.Context, filter bson.M) {
	cursor, err := Col.Find(Ctx, filter)
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(Ctx)

	for cursor.Next(Ctx) {
		var author Author
		var book Book
		if ok == true {
			if err := cursor.Decode(&author); err != nil {
				log.Fatal(err)
			}
			if err != nil {
				log.Fatal(err)
			}
			author.Book = BookList
			AuthorList = append(AuthorList, author)
		} else {
			if err := cursor.Decode(&book); err != nil {
				log.Fatal(err)
			}
			if err != nil {
				log.Fatal(err)
			}
			BookList = append(BookList, &book)
		}

	}
	if err := cursor.Err(); err != nil {
		log.Fatal(err)
	}

}

func GetAuthors(p graphql.ResolveParams) (interface{}, error) {
	filter := bson.M{}
	ok = false
	GetDataFromCollection(db.CollectionBook, db.Ctx, filter)
	defer func() {
		BookList = nil
	}()
	ok = true
	GetDataFromCollection(db.CollectionAuthor, db.Ctx, filter)
	defer func() {
		AuthorList = nil
	}()

	if len(AuthorList) == 0 {
		return nil, nil
	}
	return AuthorList, nil
}

func GetAuthorByName(p graphql.ResolveParams) (interface{}, error) {
	authorName, isOK := p.Args["author_name"].(string)
	if !isOK {
		return nil, nil
	}

	filter := bson.M{
		"authorname": authorName,
	}
	ok = true
	GetDataFromCollection(db.CollectionAuthor, db.Ctx, filter)
	defer func() {
		AuthorList = nil
	}()
	if len(AuthorList) == 0 {
		return nil, nil
	}
	return AuthorList[0], nil
}

func CreateNewAuthor(p graphql.ResolveParams) (interface{}, error) {
	filter := bson.M{}
	authorName, isOK := p.Args["author_name"].(string)
	if !isOK {
		return nil, nil
	}
	authorID := RandStringRunes(8)

	ok = false
	GetDataFromCollection(db.CollectionBook, db.Ctx, filter)
	defer func() {
		BookList = nil
	}()
	newAuthor := Author{
		ID:         authorID,
		AuthorName: authorName,
	}
	_, db.Err = db.CollectionAuthor.InsertOne(db.Ctx, newAuthor)
	if db.Err != nil {
		log.Fatal(db.Err)
	}

	fmt.Println("Author inserted successfully.")
	return newAuthor, nil
}
