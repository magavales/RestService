package auth

import (
	"github.com/golang-jwt/jwt/v4"
	"log"
	"time"
)

type Authentication struct {
	Token *jwt.Token
}

type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func (a *Authentication) GetToken(username string) (string, error) {
	jwtKey := []byte("bebra")

	claims := &Claims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "server",
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(5 * time.Minute)),
		},
	}
	a.Token = jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedString, err := a.Token.SignedString(jwtKey)
	if err != nil {
		log.Printf("Token hasn't been created! : %s\n", err)
	}

	return signedString, err
}
