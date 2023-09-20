package db

import (
	"go.mongodb.org/mongo-driver/bson"
	"graphql_test/db"
	"graphql_test/schema"
	"log"
)

func GetDataFromAuthorCollection(filter bson.M) []schema.Author {
	cursor, err := db.CollectionAuthor.Find(db.Ctx, filter)
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(db.Ctx)

	var AuthorsList []schema.Author
	var BooksList []*schema.Book
	BooksList = GetDataFromBookCollection(filter)
	for cursor.Next(db.Ctx) {
		var author schema.Author
		if err := cursor.Decode(&author); err != nil {
			log.Fatal(err)
		}
		if err != nil {
			log.Fatal(err)
		}
		author.Book = BooksList
		AuthorsList = append(AuthorsList, author)
	}
	if err := cursor.Err(); err != nil {
		log.Fatal(err)
	}
	return AuthorsList
}

func InsertAuthor(newAuthor schema.Author) error {
	_, db.Err = db.CollectionAuthor.InsertOne(db.Ctx, newAuthor)
	if db.Err != nil {
		log.Fatal(db.Err)
	}
	return nil
}
