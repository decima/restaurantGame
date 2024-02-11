package controllers

import (
	"errors"
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

	token, cookID, err := services.Container.GetRestaurantService().NewHumanCook(restaurant, newRestaurant.CookName)
	if err != nil {
		if errors.Is(err, models.ErrKitchenIsFull) {
			utils.PreconditionFailed(c, "Kitchen is full")
			return
		}
		utils.InternalServerError(c, err)
		return
	}
	cook, _ := restaurant.Kitchen.Crew.GetMember(cookID)
	utils.Created(c, output.RestaurantCreationResponse(newRestaurant.Name, cook.Name, cookID, token))

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
