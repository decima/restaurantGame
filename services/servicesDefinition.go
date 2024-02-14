package services

import (
	"restaurantAPI/lib/database"
	"restaurantAPI/models"
)

func (sc *serviceContainer) GetContainer() *serviceContainer {
	return sc.getOrPanic("container", func() (interface{}, error) {
		return sc, nil
	}).(*serviceContainer)
}

func (sc *serviceContainer) GetSecuritySigner() *securitySigner {
	return sc.getOrPanic("securitySigner", func() (interface{}, error) {
		return newSecuritySigner(), nil
	}).(*securitySigner)
}

func (sc *serviceContainer) GetAuthenticator() *authenticator {
	return sc.getOrPanic("authenticator", func() (interface{}, error) {
		return newAuthenticator(sc.GetSecuritySigner()), nil
	}).(*authenticator)
}

func (sc *serviceContainer) GetRestaurantService() *restaurantService {
	return sc.getOrPanic("restaurantService", func() (interface{}, error) {
		return newRestaurantService(sc.GetRestaurantsRepository(), sc.GetAuthenticator()), nil
	}).(*restaurantService)
}

func (sc *serviceContainer) GetRestaurantsRepository() *RestaurantsRepository {
	collection := sc.getOrPanic("restaurantsRepository", func() (interface{}, error) {
		return collection[*models.Restaurant]("restaurants")
	}).(database.Collection[*models.Restaurant])
	return (*RestaurantsRepository)(&collection)
}

func (sc *serviceContainer) GetRecipesRepository() *RecipesRepository {
	collection := sc.getOrPanic("recipesRepository", func() (interface{}, error) {
		return collection[*models.Recipe]("recipes")
	}).(database.Collection[*models.Recipe])

	return (*RecipesRepository)(&collection)
}
