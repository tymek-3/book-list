package middleware

import (
	"book-list/internal/utils"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func Auth(logger *log.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := c.Cookie("token")
		if err != nil {
			fmt.Println(c.Writer.Header())
			token := c.GetHeader("Authorization")
			fmt.Println("token")
			splitToken := strings.Split(token, "Bearer ")
			token = splitToken[1]
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
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": fmt.Sprintf("Token invalid: %e", err),
			})
			c.Abort()
			return
		}

		userIDStr, ok := claims["user_id"].(string)
		userID, err := uuid.Parse(userIDStr)
		if !ok || err != nil {
			logger.Printf("No user id or id is invalid: %e\n", err)
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "No user id or id is invalid",
			})
			c.Abort()
			return
		}

		email, ok := claims["email"].(string)
		if !ok {
			logger.Println("No user email")
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "No user email",
			})
			c.Abort()
			return
		}

		c.Set("user_id", userID.String())
		c.Set("email", email)

		c.Next()
	}
}
