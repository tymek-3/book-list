package utils

import (
	"book-list/config"
	"book-list/internal/domain/entities"
	"fmt"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var (
	ErrInvalidToken = fmt.Errorf("Invalid token")
	ErrExpiredToken = fmt.Errorf("Expired token")
)

func GenerateToken(logger *log.Logger, user entities.User) string {
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": string(user.Email),
		"role":  user.Role.Name,
		"exp":   time.Now().Add(time.Hour).Unix(),
	})

	s, err := t.SignedString([]byte(config.AppConfig.JWT_SECRET))
	if err != nil {
		logger.Fatalln(err.Error())
	}

	return s
}

func VerifyToken(logger *log.Logger, token string) (jwt.MapClaims, error) {
	p := jwt.NewParser(jwt.WithExpirationRequired())

	t, err := p.Parse(token, func(token *jwt.Token) (any, error) {
		return []byte(config.AppConfig.JWT_SECRET), nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := t.Claims.(jwt.MapClaims)
	if !ok || !t.Valid {
		return nil, ErrInvalidToken
	}

	exp, ok := claims["exp"].(float64)
	if !ok || time.Now().Unix() > int64(exp) {
		return nil, ErrExpiredToken
	}

	logger.Println("token valid")

	return claims, nil
}
