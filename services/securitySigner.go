package services

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type securitySigner struct {
}

func newSecuritySigner() *securitySigner {
	return &securitySigner{}
}

const key = "restaurant"

func (s *securitySigner) Encode(item map[string]interface{}) (string, error) {

	item["iat"] = time.Now()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims(item))
	return token.SignedString([]byte(key))
}

func (s *securitySigner) Decode(token string) (valid bool, item map[string]interface{}, err error) {
	valid = false
	item = nil
	err = nil

	t, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, nil
		}

		return []byte(key), nil
	})
	if err != nil {
		return
	}
	claims, ok := t.Claims.(jwt.MapClaims)
	if !ok || !t.Valid {
		return
	}
	valid = t.Valid
	item = map[string]interface{}(claims)
	return
}
