package auth

import (
	"book-list/internal/utils"
	"context"
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	ErrInvalidCredentials = fmt.Errorf("Credentials are incorrect")
)

func (ae *authEndpoints) LoginHandler(c *gin.Context) {
	var request LoginRequest
	err := c.BindJSON(&request)
	if err != nil {
		c.Error(err)
		c.Status(http.StatusBadRequest)
		return
	}

	token, err := ae.as.Login(c, request)
	if err != nil {
		c.Error(err)

		switch err {
		case ErrInvalidCredentials:
			c.Status(http.StatusUnauthorized) // Unauthorized?
		default:
			c.Status(http.StatusBadRequest)
		}

		return
	}

	utils.SetJwtTokenCookie(c, token)

	c.JSON(http.StatusOK, jwtTokenResponse{token})
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// returns token, error
func (as *AuthService) Login(ctx context.Context, request LoginRequest) (string, error) {
	user, err := as.db.UserGetByEmail(ctx, request.Email)
	switch err {
	case nil:
		break
	case sql.ErrNoRows:
		return "", ErrInvalidCredentials
	default:
		return "", err
	}

	passwordCorrect := utils.VerifyPassword(request.Password, user.PasswordHash)
	if !passwordCorrect {
		return "", ErrInvalidCredentials
	}

	token := utils.GenerateToken(as.logger, user.ID, user.Email)

	return token, nil
}
