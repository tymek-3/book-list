package data

import (
	"book-list/internal/domain/entities"

	"github.com/google/uuid"
)

func (u User) ToEntity() *entities.User {
	return entities.ConstructUser(
		u.ID,
		u.Name,
		u.Email,
		u.PasswordHash,
		// TODO:
		[]uuid.UUID{},
	)
}
