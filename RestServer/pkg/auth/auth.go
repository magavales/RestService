package auth

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"github.com/golang-jwt/jwt/v4"
	"log"
	"strings"
	"time"
)

const jwtKey = "bebra"

type Authentication struct {
	Token *jwt.Token
}

type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

type HeaderPayload struct {
	Alg      string `json:"alg"`
	Typ      string `json:"typ"`
	Username string `json:"username"`
	Iss      string `json:"iss"`
	Exp      int    `json:"exp"`
}

func (a *Authentication) GetToken(username string) (string, error) {

	claims := &Claims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "server",
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(5 * time.Minute)),
		},
	}
	a.Token = jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedString, err := a.Token.SignedString([]byte(jwtKey))
	if err != nil {
		log.Printf("Token hasn't been created! : %s\n", err)
	}

	return signedString, err
}

func (a *Authentication) CheckToken(token string) bool {
	token = strings.Fields(token)[1]
	tk := strings.Split(token, ".")
	data := tk[0] + "." + tk[1]
	mac := hmac.New(sha256.New, []byte(jwtKey))
	mac.Write([]byte(data))
	seg := base64.RawURLEncoding.EncodeToString(mac.Sum(nil))
	if seg == tk[2] {
		return true
	} else {
		return false
	}
}
