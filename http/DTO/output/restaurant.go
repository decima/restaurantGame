package output

import (
	"restaurantAPI/models"
	"time"
)

type restaurantCreationResponse struct {
	Name string `json:"name"`
	Cook struct {
		Name string `json:"name"`
		ID   string `json:"ID"`
	} `json:"cook"`
	Token string `json:"token"`
}

func RestaurantCreationResponse(name, cookName, cookID, token string) restaurantCreationResponse {
	return restaurantCreationResponse{
		Name: name,
		Cook: struct {
			Name string `json:"name"`
			ID   string `json:"ID"`
		}{
			Name: cookName,
			ID:   cookID,
		},
		Token: token,
	}
}

type restaurantMyResponse struct {
	Name    string                      `json:"name"`
	Kitchen RestaurantMyResponseKitchen `json:"kitchen"`
}

type RestaurantMyResponseKitchen struct {
	Equipment []string `json:"equipment"`
	Employees []string `json:"employees"`
	Inventory []string `json:"inventory"`
}

func RestaurantMyResponse(restaurant *models.Restaurant) restaurantMyResponse {
	return restaurantMyResponse{
		Name: restaurant.Name,
		Kitchen: RestaurantMyResponseKitchen{
			Equipment: restaurant.GetListOfEquipmentID(),
			Employees: restaurant.GetListOfEmployeeID(),
			Inventory: restaurant.GetListOfInventory(),
		},
	}
}

type restaurantHireTokenResponse struct {
	URL   string    `json:"url"`
	Token string    `json:"token"`
	Exp   time.Time `json:"expiresAt"`
}

func NewRestaurantHireTokenResponse(URL string, Token string, exp time.Time) restaurantHireTokenResponse {
	return restaurantHireTokenResponse{URL: URL, Token: Token, Exp: exp}

}
