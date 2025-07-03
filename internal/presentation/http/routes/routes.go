package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/imlargo/butter/api/docs"
	"github.com/imlargo/butter/internal/config"
	"github.com/imlargo/butter/internal/presentation/http/handlers"
	middlewares "github.com/imlargo/butter/internal/presentation/http/middleware"
	"github.com/imlargo/butter/internal/shared/utils"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func RegisterRoutes(router *gin.Engine) {
	router.Use(middlewares.NewCorsMiddleware(config.Config.Server.Host, []string{
		"http://localhost:5173",
		"https://app.hellobutter.io",
		"https://app-butter.vercel.app",
	}))

	registerDocs(router)

	completionsController := handlers.NewCompletionController()
	router.POST("/process", completionsController.CreateCompletion)
}

func registerDocs(router *gin.Engine) {

	host := config.Config.Server.Host
	if utils.IsLocalhostURL(host) {
		host += ":" + config.Config.Server.Port
	}

	if utils.IsHttpsURL(host) {
		docs.SwaggerInfo.Schemes = []string{"https"}
	} else {
		docs.SwaggerInfo.Schemes = []string{"http"}
	}

	docs.SwaggerInfo.Host = utils.CleanHostURL(host)
	docs.SwaggerInfo.BasePath = "/"

	schemaUrl := host
	schemaUrl += "/docs/doc.json"

	urlSwaggerJson := ginSwagger.URL(schemaUrl)
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, urlSwaggerJson))
}
