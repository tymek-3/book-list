package entities

import "github.com/google/uuid"

// Author

type Author struct {
	id       uuid.UUID
	FullName string
}

func (a Author) ID() uuid.UUID {
	return a.id
}

func ConstructAuthor(id uuid.UUID, fullName string) *Author {
	return &Author{
		id:       id,
		FullName: fullName,
	}
}

func NewAuthor(fullName string) *Author {
	id, err := uuid.NewV7()
	if err != nil {
		panic(err)
	}

	return ConstructAuthor(id, fullName)
}

// Publisher

type Publisher struct {
	id   uuid.UUID
	Name string
}

func (a Publisher) ID() uuid.UUID {
	return a.id
}

func ConstructPublisher(id uuid.UUID, name string) *Publisher {
	return &Publisher{
		id:   id,
		Name: name,
	}
}

func NewPublisher(name string) *Publisher {
	id, err := uuid.NewV7()
	if err != nil {
		panic(err)
	}

	return ConstructPublisher(id, name)
}
