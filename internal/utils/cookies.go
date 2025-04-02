package utils

import (
	"book-list/config"

	"github.com/gin-gonic/gin"
)

func SetJwtTokenCookie(c *gin.Context, token string) {
	c.SetCookie(
		config.AppConfig.TOKEN_COOKIE_NAME,
		token,
		config.AppConfig.TOKEN_COOKIE_AGE*60, // 60 minutes * 60 secs
		"/",
		"*",
		true,
		true,
	)
}

func RemoveJwtTokenCookie(c *gin.Context) {
	c.SetCookie(
		config.AppConfig.TOKEN_COOKIE_NAME,
		"",
		-1,
		"/",
		"*",
		true,
		true,
	)
}
