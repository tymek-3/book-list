package middleware

import (
	"book-list/config"
	"book-list/internal/domain/entities"
	"book-list/internal/utils"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type authOptions struct {
	RolePolicy entities.Role
}

type authOptionFunc func(*authOptions)

func WithRolePolicy(role entities.Role) authOptionFunc {
	return func(o *authOptions) {
		o.RolePolicy = role
	}
}

func Auth(logger *log.Logger, funcOpts ...authOptionFunc) gin.HandlerFunc {
	options := authOptions{
		RolePolicy: entities.RoleUser,
	}

	for _, opt := range funcOpts {
		opt(&options)
	}

	return func(c *gin.Context) {
		token, err := c.Cookie(config.AppConfig.TOKEN_COOKIE_NAME)
		if err != nil {
			// TODO: fix
			fmt.Print("header: ")
			fmt.Println(c.Writer.Header())
			token := c.GetHeader("Authorization")
			fmt.Printf("auth header: %s\n", token)
			if token != "" {
				splitToken := strings.Split(token, "Bearer ")
				token = splitToken[1]
			}
		}

		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "No token",
			})
			c.Abort()
			return
		}

		claims, err := utils.VerifyToken(logger, token)
		if err != nil {
			msg := fmt.Sprintf("Token invalid: %e", err)
			logger.Println(msg)
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": msg,
			})
			c.Abort()
			return
		}

		email, ok := claims["email"].(string)
		if !ok {
			msg := "No user email"
			logger.Println(msg)
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": msg,
			})
			c.Abort()
			return
		}

		// TODO: split authentication and authorization
		roleStr, ok := claims["role"].(string)
		var role entities.Role
		if !ok {
			logger.Println("role not in claims")
			role = entities.RoleUser
		} else {
			role = entities.RoleFromString(roleStr)
		}

		if role.Level < options.RolePolicy.Level {
			msg := "Role not authorized"
			logger.Printf("%s, %s trying to access %s", msg, role.Name, options.RolePolicy.Name)
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": msg,
			})
			c.Abort()
			return
		}

		c.Set("email", entities.Email(email))
		c.Set("role", role)

		c.Next()
	}
}
