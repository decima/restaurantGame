package services

import (
	"fmt"
	"github.com/google/uuid"
	"restaurantAPI/lib/database"
	"restaurantAPI/lib/faker"
	"restaurantAPI/models"
	"time"
)

type restaurantService struct {
	repository     *RestaurantsRepository
	authenticator  *authenticator
	securitySigner *securitySigner
}

func newRestaurantService(
	repository *RestaurantsRepository,
	authenticator *authenticator,
	securitySigner *securitySigner,
) *restaurantService {
	return &restaurantService{
		repository:     repository,
		authenticator:  authenticator,
		securitySigner: securitySigner,
	}
}

func (rs *restaurantService) NewRestaurant(name string, email *string) (restaurant *models.Restaurant, err error) {

	restaurant = models.NewRestaurant(name, email)

	//get first cook
	id, err := uuid.NewV6()
	if err != nil {
		restaurant = nil
		return
	}
	restaurant.ID = id.String()

	err = (*rs.repository).Insert(restaurant)
	if err != nil {
		restaurant = nil
		return
	}
	return
}

func (rs *restaurantService) NewHumanCook(restaurant *models.Restaurant, cookName *string, role string) (token string, employeeID string, err error) {
	employeeID = ""
	token = ""
	err = nil
	cookNameValue := cookNameToString(cookName)
	employee := models.NewCrewMate(cookNameValue, role)
	err = restaurant.HireEmployee(employee)
	employeeID = employee.ID
	if err != nil {
		return
	}
	err = (*rs.repository).Update(restaurant)
	if err != nil {
		return
	}
	token, err = rs.authenticator.Register(restaurant, employee.ID)
	if err != nil {
		return
	}
	return
}

func cookNameToString(cookName *string) string {
	cookNameValue := faker.PersonName()
	if cookName != nil {
		cookNameValue = *cookName
	}
	return cookNameValue
}

func (rs *restaurantService) Close(restaurant *models.Restaurant) error {
	return (*rs.repository).Delete(restaurant.GetID())
}

func (rs *restaurantService) CreateHireToken(restaurant *models.Restaurant) (string, time.Time, error) {
	if ok, err := restaurant.CanHire(); !ok {
		return "", time.Now(), err
	}
	expiration := time.Now().Add(time.Minute * 24)
	token, err := rs.securitySigner.Crypt(map[string]interface{}{
		"restaurant_id":   restaurant.GetID(),
		"action":          "hire",
		"currentCrewSize": len(restaurant.Kitchen.Crew),
		"exp":             expiration,
	})

	return token, expiration, err
}

func (rs *restaurantService) ValidateJoinDemand(restaurantID string, token string, cookName *string) (restaurant *models.Restaurant, valid bool, err error) {
	restaurant, err = (*rs.repository).Find(database.ID(restaurantID))
	valid = false
	if err != nil {
		return
	}
	valid, content, err := rs.securitySigner.Decrypt(token)
	if !valid {
		err = fmt.Errorf("invalid token")
		return
	}
	if err != nil {
		return
	}
	if content["restaurant_id"] != restaurantID {
		err = fmt.Errorf("invalid restaurant")
		return
	}
	if content["action"] != "hire" {
		err = fmt.Errorf("invalid action")
		return
	}
	if content["exp"].(time.Time).Before(time.Now()) {
		err = fmt.Errorf("token expired")
		return

	}
	if content["currentCrewSize"].(int) != len(restaurant.Kitchen.Crew) {
		err = fmt.Errorf("invalid crew size")
		return

	}
	valid = true
	return

}
