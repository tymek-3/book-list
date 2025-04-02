package auth

import (
	"book-list/internal/data"
	"book-list/internal/middleware"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthService struct {
	logger *log.Logger
	db     *data.Queries
}

type authEndpoints struct {
	as *AuthService
	r  *gin.RouterGroup
}

func AddAuth(router *gin.RouterGroup, logger *log.Logger, db *data.Queries) {
	r := router.Group("/auth")

	aService := &AuthService{logger, db}
	aEndoints := &authEndpoints{aService, r}

	r.POST("/signup", aEndoints.SignUpHandler)
	r.POST("/login", aEndoints.LoginHandler)
	r.POST("/logout", middleware.Auth(logger), aEndoints.LogoutHandler)

	// TODO: remove later
	r.GET("/test", func(c *gin.Context) {
		email := c.Query("email")
		_, err := db.UserGetByEmail(c, email)
		if err != nil {
			c.String(http.StatusNotFound, err.Error())
		}
		c.String(http.StatusOK, "exists")
	})
}
