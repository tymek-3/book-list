package auth

import (
	"book-list/internal/data"
	"book-list/internal/domain/entities"
	"book-list/internal/utils"
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	ErrDuplicateEmail = fmt.Errorf("User with this email already exists.")
)

func (ae *authEndpoints) SignUpHandler(c *gin.Context) {
	var request SignUpRequest
	err := c.BindJSON(&request)
	if err != nil {
		c.Error(err)
		c.Status(http.StatusBadRequest)
		return
	}

	token, err := ae.as.SignUp(c, request)
	if err != nil {
		c.Error(err)
		c.Status(http.StatusBadRequest)
		return
	}

	c.JSON(http.StatusCreated, jwtTokenResponse{token})

	utils.SetJwtTokenCookie(c, token)
}

type SignUpRequest struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// returns jwtToken, err
func (as *AuthService) SignUp(ctx context.Context, request SignUpRequest) (string, error) {
	_, err := as.db.UserGetByEmail(ctx, request.Email)
	if err == nil {
		return "", ErrDuplicateEmail
	}

	pHash, err := utils.HashPassword(request.Password)
	if err != nil {
		return "", err
	}

	user, err := entities.NewUser(request.Name, request.Email, pHash)
	if err != nil {
		return "", err
	}

	as.db.UserAdd(ctx, data.UserAddParams{
		ID:           user.ID(),
		Name:         user.Name,
		Email:        user.Email,
		PasswordHash: user.PasswordHash,
	})

	token := utils.GenerateToken(as.logger, user.ID(), user.Email)
	return token, nil
}
