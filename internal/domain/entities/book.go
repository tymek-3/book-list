package entities

import (
	"time"

	"github.com/google/uuid"
)

type Book struct {
	id              uuid.UUID
	Name            string
	Score           float32
	PublicationDate time.Time
	Type            BookType
	Author          Author
	Publisher       Publisher
}

func (b Book) ID() uuid.UUID {
	return b.id
}

func ConstructBook(id uuid.UUID, name string, score float32, publicationDate time.Time, bookType BookType, author Author, publisher Publisher) *Book {
	return &Book{
		id:              id,
		Name:            name,
		Score:           score,
		PublicationDate: publicationDate,
		Type:            bookType,
		Author:          author,
		Publisher:       publisher,
	}
}

func NewBook(name string, publicationDate time.Time, bookType BookType, author Author, publisher Publisher) *Book {
	id, err := uuid.NewV7()
	if err != nil {
		panic(err)
	}

	return ConstructBook(id, name, 0.0, publicationDate, bookType, author, publisher)
}

type BookType struct {
	id   uuid.UUID
	Name string
}

func (b BookType) ID() uuid.UUID {
	return b.id
}

func ConstructBookType(id uuid.UUID, name string) *BookType {
	return &BookType{
		id:   id,
		Name: name,
	}
}

func NewBookType(name string) *BookType {
	id, err := uuid.NewV7()
	if err != nil {
		panic(err)
	}

	return ConstructBookType(id, name)
}
