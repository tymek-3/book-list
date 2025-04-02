package entities

import "github.com/google/uuid"

type Role string

var (
	RoleUser  Role = "user"
	RoleAdmin Role = "admin"
)

type User struct {
	id           uuid.UUID
	Name         string
	Email        string
	Role         Role
	PasswordHash string
	List         []uuid.UUID
}

func (u User) ID() uuid.UUID {
	return u.id
}

func ConstructUser(id uuid.UUID, name, email, passwordHash string, role Role, list []uuid.UUID) *User {
	return &User{
		id:           id,
		Name:         name,
		Email:        email,
		PasswordHash: passwordHash,
		List:         []uuid.UUID{},
	}
}

type UserOption func(*User)

func WithRole(role Role) UserOption {
	return func(u *User) {
		u.Role = role
	}
}

func NewUser(name, email, passwordHash string, opts ...UserOption) (*User, error) {
	id, err := uuid.NewV7()
	if err != nil {
		panic(err)
	}

	// TODO: validation

	user := ConstructUser(id, name, email, passwordHash, RoleUser, []uuid.UUID{})

	for _, o := range opts {
		o(user)
	}

	return user, nil
}
