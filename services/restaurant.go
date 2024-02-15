package services

import (
	"github.com/google/uuid"
	"restaurantAPI/lib/faker"
	"restaurantAPI/models"
)

type restaurantService struct {
	repository    *RestaurantsRepository
	authenticator *authenticator
}

func newRestaurantService(
	repository *RestaurantsRepository,
	authenticator *authenticator,
) *restaurantService {
	return &restaurantService{
		repository:    repository,
		authenticator: authenticator,
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

func (rs *restaurantService) NewHumanCook(restaurant *models.Restaurant, cookName *string) (token string, employeeID string, err error) {
	employeeID = ""
	token = ""
	err = nil
	cookNameValue := faker.PersonName()
	if cookName != nil {
		cookNameValue = *cookName
	}
	employee := models.NewCrewMate(cookNameValue)
	err = restaurant.HireEmployee(employee)
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

	employeeID = employee.ID
	return
}

func (rs *restaurantService) Close(restaurant *models.Restaurant) error {
	return (*rs.repository).Delete(restaurant.GetID())
}
