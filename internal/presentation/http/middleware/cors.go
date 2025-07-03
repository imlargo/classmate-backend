package middlewares

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func NewCorsMiddleware(host string, origins []string) gin.HandlerFunc {

	config := cors.Config{
		AllowMethods: []string{
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
			http.MethodPatch,
			http.MethodDelete,
			http.MethodHead,
			http.MethodOptions,
		},
		AllowCredentials: true,
		AllowHeaders: []string{
			"Origin",
			"Accept",
			"Authorization",
			"Content-Type",
		},
		AllowAllOrigins: true,
	}

	return cors.New(config)
}
