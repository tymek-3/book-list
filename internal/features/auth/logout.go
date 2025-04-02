package auth

import (
	"book-list/internal/utils"

	"github.com/gin-gonic/gin"
)

func (ae *authEndpoints) LogoutHandler(c *gin.Context) {
	utils.RemoveJwtTokenCookie(c)
}
