package admin

import (
	"book-list/internal/data"
	"book-list/internal/domain/entities"
	"context"
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (ae *adminEndpoints) PromoteUserHandler(c *gin.Context) {
	var request promoteUserRequest
	err := c.BindJSON(&request)
	if err != nil {
		c.Error(err)
		return
	}

	err = ae.as.PromoteUser(c, request)
	if err != nil {
		c.Error(err)
		return
	}

	c.Status(http.StatusOK)
}

type promoteUserRequest struct {
	Email string `json:"email" binding:"required"`
	Role  string `json:"role" binding:"required"`
}

var (
	ErrUserNotFound = fmt.Errorf("User not found.")
	ErrRoleToLow    = fmt.Errorf("You cannot promote user to role higher than yours.")
)

func (as *AdminService) PromoteUser(ctx context.Context, request promoteUserRequest) error {
	promoteToRole := entities.RoleFromString(request.Role)
	userRole := ctx.Value("role").(entities.Role)
	userEmail := ctx.Value("email").(entities.Email)

	if promoteToRole.Level > userRole.Level {
		return ErrRoleToLow
	}

	u, err := as.db.UserGetByEmail(ctx, request.Email)
	switch err {
	case nil:
		break
	default:
		return err
	case sql.ErrNoRows:
		return ErrUserNotFound
	}

	err = as.db.UserUpdate(ctx, data.UserUpdateParams{
		Email:        u.Email,
		Name:         u.Name,
		PasswordHash: u.PasswordHash,

		Email_2: request.Email,
		Role:    promoteToRole.Name,
		RoleSetBy: sql.NullString{
			String: string(userEmail),
			Valid:  true,
		},
	})
	if err != nil {
		return err
	}

	return nil
}
