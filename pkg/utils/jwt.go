package utils

import (
	"dropx/pkg/config"
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type Claims struct {
	UserID string `json:"user_id"`
	Email  string `json:"email"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

type ResetPasswordClaims struct {
	Email   string `json:"email"`
	Purpose string `json:"purpose"`
	jwt.RegisteredClaims
}

func GenerateToken(userID, mail, role string) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		UserID: userID,
		Email:  mail,
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.Global.JwtSecret))
}

func VerifyToken(tokenStr string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(
		tokenStr,
		&Claims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(config.Global.JwtSecret), nil
		},
	)
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return nil, jwt.ErrSignatureInvalid
	}
	return claims, nil
}

func GenerateResetPasswordToken(email string) (string, error) {
	claims := &ResetPasswordClaims{
		Email:   email,
		Purpose: "reset_password",
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(5 * time.Minute)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.Global.JwtSecret))
}

func VerifyResetPasswordToken(tokenStr string) (*ResetPasswordClaims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &ResetPasswordClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.Global.JwtSecret), nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*ResetPasswordClaims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid or expired token")
	}
	return claims, nil
}
