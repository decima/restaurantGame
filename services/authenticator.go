package services

import (
	"fmt"
	"restaurantAPI/models"
)

type authenticator struct {
	securitySigner *securitySigner
}

func newAuthenticator(securitySigner *securitySigner) *authenticator {
	return &authenticator{securitySigner: securitySigner}
}

func (a *authenticator) Register(restaurant *models.Restaurant, cookID string) (string, error) {
	return a.securitySigner.Crypt(map[string]interface{}{
		"rid": restaurant.ID,
		"cid": cookID,
	})
}

func (a *authenticator) Verify(token string) (valid bool, restaurantID string, cookID string, err error) {
	valid = false
	restaurantID = ""
	cookID = ""
	err = nil

	valid, claims, err := a.securitySigner.Decrypt(token)
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
