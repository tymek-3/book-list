package auth

import (
	"book-list/internal/data"
	"book-list/internal/domain/entities"
	"book-list/internal/utils"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

type signUpRequest struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (ae *authEndpoints) SignUpHandler(c *gin.Context) {
	var request signUpRequest
	err := c.BindJSON(&request)
	if err != nil {
		c.Error(err)
		c.Status(http.StatusBadRequest)
		return
	}

	token, err := ae.as.SignUp(c, request.Name, request.Email, request.Password)
	c.JSON(http.StatusCreated, gin.H{
		"token": token,
	})
}

// returns jwtToken, err
func (as *AuthService) SignUp(ctx context.Context, name, email, password string) (string, error) {
	pHash, err := utils.HashPassword(password)
	if err != nil {
		return "", err
	}

	user, err := entities.NewUser(name, email, pHash)
	if err != nil {
		return "", err
	}

	as.db.AddUser(ctx, data.AddUserParams{
		ID:           user.ID(),
		Name:         user.Name,
		Email:        user.Email,
		PasswordHash: user.PasswordHash,
	})

	token := utils.GenerateToken(as.logger, user.ID(), user.Email)
	return token, nil
}
