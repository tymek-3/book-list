package entities

import "github.com/google/uuid"

type User struct {
	id           uuid.UUID
	Name         string
	Email        string
	PasswordHash string
	List         []uuid.UUID
}

func (u User) ID() uuid.UUID {
	return u.id
}

func ConstructUser(id uuid.UUID, name, email, passwordHash string, list []uuid.UUID) *User {
	return &User{
		id:           id,
		Name:         name,
		Email:        email,
		PasswordHash: passwordHash,
		List:         []uuid.UUID{},
	}
}

func NewUser(name, email, passwordHash string) (*User, error) {
	id, err := uuid.NewV7()
	if err != nil {
		panic(err)
	}

	// TODO: validation

	return ConstructUser(id, name, email, passwordHash, []uuid.UUID{}), nil
}
