package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (ae *authEndpoints) LoginHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"email":   c.GetString("email"),
		"user_id": c.GetString("user_id"),
	})
}
