package db

import (
	"go.mongodb.org/mongo-driver/bson"
	"graphql_test/db"
	"graphql_test/schema"
	"log"
)

func GetDataFromBookCollection(filter bson.M) []*schema.Book {
	cursor, err := db.CollectionBook.Find(db.Ctx, filter)
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(db.Ctx)

	var BooksList []*schema.Book

	for cursor.Next(db.Ctx) {
		var book schema.Book
		if err := cursor.Decode(&book); err != nil {
			log.Fatal(err)
		}
		if err != nil {
			log.Fatal(err)
		}
		BooksList = append(BooksList, &book)
	}
	if err := cursor.Err(); err != nil {
		log.Fatal(err)
	}
	return BooksList
}

func InsertBook(newBook *schema.Book) error {
	_, db.Err = db.CollectionBook.InsertOne(db.Ctx, newBook)
	if db.Err != nil {
		log.Fatal(db.Err)
	}
	return nil
}
