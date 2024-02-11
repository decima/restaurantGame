package middlewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"restaurantAPI/http/utils"
	"restaurantAPI/lib/database"
	"restaurantAPI/services"
	"strings"
)

func AuthenticatorMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			utils.Unauthorized(c)
			return
		}
		token = strings.Split(token, " ")[1]
		valid, restaurantID, cookID, err := services.Container.GetAuthenticator().Verify(token)

		if err != nil {
			utils.Unauthorized(c)
			log.Debug().Err(err).Msg("Error verifying token")
			return
		}
		if !valid {
			utils.Unauthorized(c)
			log.Debug().Msg("Invalid token")
			return
		}
		repository := services.Container.GetRestaurantsRepository()

		restaurant, err := (*repository).Find(database.ID(restaurantID))

		if err != nil {
			utils.Unauthorized(c)
			log.Debug().Err(err).Msg("Error finding restaurant")
			return
		}
		utils.SetRestaurant(c, restaurant)

		if cookID != "" {
			fmt.Println(cookID)
			employee, ok := restaurant.Kitchen.Crew.GetMember(cookID)

			if !ok {
				utils.Unauthorized(c)
				log.
					Debug().
					Msg("Crew not found")
				return
			}
			if employee.Bot {
				utils.Unauthorized(c)
				log.
					Debug().
					Msg("Bot not allowed")
				return
			}
			utils.SetCook(c, employee)

		}

		c.Next()
	}
}
