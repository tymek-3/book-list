package auth

import (
	"book-list/internal/data"
	"book-list/internal/middleware"
	"log"

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
	aEndoints := authEndpoints{aService, r}

	r.POST("/signup", aEndoints.SignUpHandler)
	r.POST("/login", middleware.Auth(logger), aEndoints.LoginHandler)
}
