package service

import (
	"time"

	"github.com/chaiyapatoam/go-fiber-crud/config"
	"github.com/golang-jwt/jwt/v5"
)

type TokenMetadata struct {
	Email   string
	Expires int64
}

func GenerateToken(email string) (string, error) {
	secret := config.Config("JWT_SECRET")

	// Claims is like payload
	claims := jwt.MapClaims{}
	claims["email"] = email
	claims["exp"] = time.Now().Add(time.Minute * time.Duration(60*24)).Unix() //(60 mins) * 24

	// Create Token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign with Secret
	t, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return t, nil
}

func ValidateToken(token string) (*TokenMetadata, error) {
	decodedToken, err := jwt.Parse(token, jwtKeyFunc)
	if err != nil {
		return nil, err
	}
	claims, ok := decodedToken.Claims.(jwt.MapClaims)
	if ok && decodedToken.Valid {
		expires := int64(claims["exp"].(float64))
		email := claims["email"].(string)

		return &TokenMetadata{
			Email:   email,
			Expires: expires,
		}, nil
	}
	return nil, err

	// fmt.Println(decodedToken)
}

func jwtKeyFunc(token *jwt.Token) (interface{}, error) {
	secret := config.Config("JWT_SECRET")

	return []byte(secret), nil
}
