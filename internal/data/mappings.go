package data

import (
	"book-list/internal/domain/entities"

	"github.com/google/uuid"
)

func (u User) ToEntity() *entities.User {
	var role entities.Role
	switch u.Role {
	default:
	case "user":
		role = entities.RoleUser
		break
	case "admin":
		role = entities.RoleAdmin
		break
	}

	return entities.ConstructUser(
		u.ID,
		u.Name,
		u.Email,
		u.PasswordHash,
		role,
		// TODO:
		[]uuid.UUID{},
	)
}
