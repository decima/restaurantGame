package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"strings"
)

func Error(c *gin.Context, code int, defaultMessage any, message ...string) {
	if len(message) > 0 {
		Error(c, code, strings.Join(message, ", "))
		return
	}
	log.Debug().Msgf("Error %d: %v", code, defaultMessage)

	Response(c, code, gin.H{"error": defaultMessage})
}

func BadRequest(c *gin.Context, message ...string) {
	Error(c, 400, "Bad Request", message...)
}

func Unauthorized(c *gin.Context, message ...string) {
	Error(c, 401, "Unauthorized", message...)
}

func Forbidden(c *gin.Context, message ...string) {
	Error(c, 403, "Resource Forbidden", message...)
}

func NotFound(c *gin.Context, message ...string) {
	Error(c, 404, "Resource Not found", message...)
}
func PreconditionFailed(c *gin.Context, message ...string) {
	Error(c, 412, "Precondition Failed", message...)
}

func InternalServerError(c *gin.Context, err error) {
	Error(c, 500, "Internal Server Error", err.Error())
	log.Err(err).Msg("Internal Server Error")
}

func Response(c *gin.Context, code int, content any) {
	c.AbortWithStatusJSON(code, content)
}

func Ok(c *gin.Context, content any) {
	Response(c, 200, content)
}

func Created(c *gin.Context, content any) {
	Response(c, 201, content)
}

func NoContent(c *gin.Context) {
	Response(c, 204, nil)
}
