package jwt

import (
	"github.com/dgrijalva/jwt-go"
	jsoniter "github.com/json-iterator/go"
)

const SECRET = "UJna5NUrKdinwcZekpqTastdSWfE5xsa"

func CreateToken(jwtMap jwt.MapClaims) (string, error) {
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtMap)
	token, err := at.SignedString([]byte(SECRET))
	if err != nil {
		return "", err
	}
	return token, nil
}

func ParseToken(token string) (string, error) {
	claim, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(SECRET), nil
	})
	if err != nil {
		return "", err
	}
	re, err := jsoniter.Marshal(claim.Claims.(jwt.MapClaims))
	if err != nil {
		return "", err
	}
	return string(re), nil
}
