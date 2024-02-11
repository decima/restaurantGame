package services

import (
	"fmt"
	"restaurantAPI/config"
	"restaurantAPI/lib/database"
	"restaurantAPI/lib/database/mongoAdapter"
	"restaurantAPI/models"
)

func collection[T database.Entity](name database.CollectionName) (database.Collection[T], error) {
	db, err := getDatabaseClient()
	if err != nil {
		return nil, err
	}

	switch config.DatabaseURI().Scheme {
	case "mongodb":
		db2 := db.(*mongoAdapter.Client)
		collection := mongoAdapter.GetCollection[T](db2, name)
		return &collection, nil
	}
	return nil, fmt.Errorf("unsupported database engine: %s", config.DatabaseURI().Scheme)
}

type RestaurantsRepository database.Collection[*models.Restaurant]
type RecipesRepository database.Collection[*models.Recipe]

var databaseInstance database.Client

func getDatabaseClient() (database.Client, error) {

	dbURI := config.DatabaseURI()
	var err error

	if databaseInstance != nil {
		return databaseInstance, nil
	}

	switch dbURI.Scheme {
	case "mongodb":
		databaseInstance = mongoAdapter.NewClient(dbURI)

	default:
		return nil, fmt.Errorf("unsupported database engine: %s", dbURI.Scheme)
	}

	return databaseInstance, err

}
