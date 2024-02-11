package services

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/rs/zerolog/log"
	"restaurantAPI/models"
	"time"
)

type authenticator struct {
}

func newAuthenticator() *authenticator {
	return &authenticator{}
}

const key = "restaurant"

func (a *authenticator) Register(restaurant *models.Restaurant, cookID string) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"rid": restaurant.ID,
		"cid": cookID,
		"iat": time.Now(),
	})
	//using constant key
	return token.SignedString([]byte(key))
}

func (a *authenticator) Verify(token string) (valid bool, restaurantID string, cookID string, err error) {
	valid = false
	restaurantID = ""
	cookID = ""
	err = nil

	t, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(key), nil
	})
	if err != nil {
		log.Debug().Err(err).Msg("Error parsing token")
		return
	}
	claims, ok := t.Claims.(jwt.MapClaims)
	if !ok || !t.Valid {
		err = fmt.Errorf("invalid claims")
		return
	}
	valid = t.Valid
	rid, ok := claims["rid"].(string)
	if !ok {
		err = fmt.Errorf("rid not found")
		return
	}
	restaurantID = rid

	cid, ok := claims["cid"].(string)
	if !ok {
		err = fmt.Errorf("cid not found")
		return
	}
	cookID = cid
	return
}
