package main

import (
	"github.com/gin-gonic/gin"
	"github.com/imlargo/butter/internal/config"
	"github.com/imlargo/butter/internal/presentation/http/routes"
	"github.com/imlargo/butter/internal/shared/utils"
)

// @title Butter API
// @version 1.0
// @description Backend service for Butter

// @contact.name Default
// @contact.url https://default.dev
// @license.name MIT
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization

// @securityDefinitions.apiKey PushApiKey
// @in header
// @name X-API-Key
func main() {
	config.LoadConfig()

	router := gin.Default()

	routes.RegisterRoutes(router)

	addr := utils.CleanHostURL(config.Config.Server.Host + ":" + config.Config.Server.Port)
	router.Run(addr)
}
