package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"restaurantAPI/http/DTO/input"
	"restaurantAPI/http/DTO/output"
	"restaurantAPI/http/utils"
	"restaurantAPI/models"
	"restaurantAPI/services"
)

type RestaurantController struct{}

func NewRestaurantController() *RestaurantController {
	return &RestaurantController{}
}

func (rc *RestaurantController) NewRestaurant(c *gin.Context) {
	if utils.GetRestaurant(c) != nil {
		utils.BadRequest(c, "Can't create a new restaurant while logged in")
		return
	}

	var newRestaurant input.RestaurantCreation
	if err := c.BindJSON(&newRestaurant); err != nil {
		utils.BadRequest(c)
		return
	}

	restaurant, err := services.Container.GetRestaurantService().NewRestaurant(newRestaurant.Name, newRestaurant.Email)

	if err != nil {
		utils.InternalServerError(c, err)
		return
	}

	token, cookID, _ := services.Container.GetRestaurantService().NewHumanCook(restaurant, newRestaurant.CookName, models.Owner)
	fmt.Println(cookID)
	cook, _ := restaurant.Kitchen.Crew.GetMember(cookID)

	fmt.Println(cook.Name)
	fmt.Println(cookID)
	fmt.Println(token)
	utils.Created(c, output.RestaurantCreationResponse(restaurant.Name, cook.Name, cookID, token))

}

func (rc *RestaurantController) GetHireToken(c *gin.Context) {
	restaurant := utils.GetRestaurantOrAbort(c)
	if restaurant == nil {
		return
	}

	if !utils.IsOwner(c) {
		utils.Forbidden(c)
		return
	}

	token, exp, err := services.Container.GetRestaurantService().CreateHireToken(restaurant)
	if err != nil {
		utils.InternalServerError(c, err)
		return
	}
	url := *c.Request.URL
	url.Path = fmt.Sprintf("/restaurant/%s/join", restaurant.ID)

	utils.Ok(c, output.NewRestaurantHireTokenResponse(url.String(), token, exp))
}

func (rc *RestaurantController) Join(c *gin.Context) {
	restaurant := utils.GetRestaurant(c)
	if restaurant != nil {
		utils.BadRequest(c, "Can't join a restaurant while logged in")
		return
	}

	restaurantID := c.Param("restaurant")
	payload := input.RestaurantJoin{}
	if err := c.BindJSON(&payload); err != nil {
		utils.BadRequest(c)
		return
	}
	restaurant, validDemand, err := services.Container.GetRestaurantService().ValidateJoinDemand(restaurantID, payload.Token, payload.CookName)

	if !validDemand {
		utils.BadRequest(c, err.Error())
		return
	}

	token, cookID, err := services.Container.GetRestaurantService().NewHumanCook(restaurant, payload.CookName, models.Cook)
	if err != nil {
		utils.BadRequest(c, err.Error())
		return
	}
	member, _ := restaurant.Kitchen.Crew.GetMember(cookID)
	utils.Ok(c, output.RestaurantCreationResponse(restaurant.Name, member.Name, cookID, token))
}

func (rc *RestaurantController) MyRestaurant(c *gin.Context) {
	restaurant := utils.GetRestaurant(c)
	if restaurant == nil {
		utils.NotFound(c)
		return
	}

	utils.Ok(c, output.RestaurantMyResponse(restaurant))
}

func (rc *RestaurantController) Close(c *gin.Context) {
	restaurant := utils.GetRestaurant(c)
	if restaurant == nil {
		utils.NotFound(c)
		return
	}

	force := c.Request.URL.Query().Has("force")
	if !force {
		utils.BadRequest(c, "You must force to close your restaurant by passing the query parameter 'force'")
		return
	}

	services.Container.GetRestaurantService().Close(restaurant)
	c.JSON(204, nil)
}
