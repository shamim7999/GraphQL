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

func GetBooksByAuthID(p graphql.ResolveParams) (interface{}, error) {
	authorID, isOK := p.Args["author_id"].(string)
	if !isOK {
		return nil, nil
	}

	var matchingAuthors []schema.Author
	var boi []schema.Book

	for _, ids := range schema.AuthorList {
		if ids.ID == authorID {
			matchingAuthors = append(matchingAuthors, ids)
		}
	}
	boi = append(boi, schema.Book{
		Author:    matchingAuthors,
		Author_Id: authorID,
	})
	return boi, nil
}

func GetBooks(p graphql.ResolveParams) (interface{}, error) {
	return schema.BookList, nil
}

func CreateNewAuthor(params graphql.ResolveParams) (interface{}, error) {

	newID := RandStringRunes(8)

	var (
		texts   string
		authors string
	)

	//fmt.Printf("Enter a Author Name: ")

	//_, err := fmt.Scanf("%v\n", &authors)
	//if err != nil {
	//	fmt.Printf("Error: ", err)
	//	return nil, nil
	//}

	//fmt.Printf("Enter some Texts: ")

	//_, err = fmt.Scanf("%v\n", &texts)
	//if err != nil {
	//	fmt.Printf("Error: ", err)
	//	return nil, nil
	//}
	authors = "Shamim Sarker"
	texts = "Clean Architecture"
	newAuthor := schema.Author{
		ID:      newID,
		Text:    texts,
		Authors: authors,
	}

	schema.AuthorList = append(schema.AuthorList, newAuthor)
	return newAuthor, nil
}
