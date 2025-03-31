package utils

import (
	"book-list/config"
	"fmt"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

var (
	ErrInvalidToken = fmt.Errorf("Invalid token")
	ErrExpiredToken = fmt.Errorf("Expired token")
)

func GenerateToken(logger *log.Logger, userID uuid.UUID) string {
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID.String(),
		"exp":     time.Now().Add(time.Hour).Unix(),
	})

	s, err := t.SignedString([]byte(config.AppConfig.JWT_SECRET))
	if err != nil {
		logger.Fatalln(err.Error())
	}

	return s
}

func VerifyToken(logger *log.Logger, token string) (jwt.MapClaims, error) {
	p := jwt.NewParser(jwt.WithExpirationRequired())

	t, err := p.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.AppConfig.JWT_SECRET), nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := t.Claims.(jwt.MapClaims)
	if !ok || !t.Valid {
		return nil, ErrInvalidToken
	}

	exp, ok := claims["exp"].(int64)
	if !ok || time.Now().Unix() > exp {
		return nil, ErrExpiredToken
	}

	logger.Println("token valid")

	return claims, nil
}
