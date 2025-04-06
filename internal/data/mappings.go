package data

import (
	"book-list/internal/domain/entities"

	"github.com/google/uuid"
)

func (u User) ToEntity() *entities.User {
	return entities.ConstructUser(
		entities.Email(u.Email),
		u.Name,
		u.PasswordHash,
		entities.RoleFromString(u.Role),
		// TODO:
		[]uuid.UUID{},
	)
}
