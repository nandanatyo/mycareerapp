package jwt

import (
	"errors"
	"mycareerapp/internal/infra/env"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type JWTI interface {
	GenerateToken(userID uuid.UUID, isAdmin bool) (string, error)
	ValidateToken(token string) (uuid.UUID, bool, error)
}

type JWT struct {
	secretKey   string
	expiredTime int
}

func NewJwt(env *env.Env) JWTI {
	secretKey := env.JwtSecret
	expiredTime := env.JwtExpired

	return &JWT{
		secretKey:   secretKey,
		expiredTime: expiredTime,
	}
}

type Claims struct {
	ID      uuid.UUID
	IsAdmin bool
	jwt.RegisteredClaims
}

func (j *JWT) GenerateToken(userID uuid.UUID, isAdmin bool) (string, error) {
	claim := Claims{
		ID:      userID,
		IsAdmin: isAdmin,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * time.Duration(j.expiredTime))),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	tokenString, err := token.SignedString([]byte(j.secretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (j *JWT) ValidateToken(tokenString string) (uuid.UUID, bool, error) {
	var claim Claims

	token, err := jwt.ParseWithClaims(tokenString, &claim, func(token *jwt.Token) (interface{}, error) {
		return []byte(j.secretKey), nil
	})

	if err != nil {
		return uuid.Nil, false, err
	}

	if !token.Valid {
		return uuid.Nil, false, errors.New("token invalid")
	}

	userID := claim.ID
	isAdmin := claim.IsAdmin

	return userID, isAdmin, nil
}
