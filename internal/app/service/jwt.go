package service

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var jwtSecretKey = []byte("secret")

type JWTActions interface {
	GenerateToken(userID uint) (string, error)
	VerifyToken(tokenString string) (*Claims, error)
}

type JWTService struct {
	JWTActions JWTActions
}

func NewJWTService() *JWTService {
	return &JWTService{}
}

type Claims struct {
	UserID uint `json:"user_id"`
	jwt.RegisteredClaims
}

// GenerateToken implements handler.JWTActions.
func (*JWTService) GenerateToken(userID uint) (string, error) {
	expirationTime := time.Now().Add(1 * time.Hour)

	claims := &Claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtSecretKey)

	return tokenString, err
}

// VerifyToken implements handler.JWTActions.
func (*JWTService) VerifyToken(tokenString string) (*Claims, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtSecretKey, nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, jwt.ErrSignatureInvalid
	}

	return claims, nil
}
