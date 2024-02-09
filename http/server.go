package http

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"restaurantAPI/http/router"
)

func Serve(address string) error {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	v1 := r.Group("/api/v1")
	// Routes
	router.RecipeRoutes(v1)

	log.Info().Str("domain", "server").Msg("Starting server on " + address)
	return r.Run(address)
}
