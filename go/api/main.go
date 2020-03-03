package main

import (
	"fmt"
	"net/http"
	"randokiak-api/controller"

	_ "randokiak-api/docs"

	"rdk.io/utils"

	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	ginSwagger "github.com/swaggo/gin-swagger"
)

const (
	defaultPort = "8082"
)

// @title Swagger randokiakapi
// @version 1.0
// @description This is public swagger for randokiakapi project.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @BasePath /api/v1

// @securityDefinitions.basic BasicAuth

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

func main() {
	r := gin.Default()
	c := controller.New()

	v1 := r.Group("/api/v1")
	{
		api := v1.Group("/profiles")
		{
			api.GET("", c.AskMoreProfiles)
		}
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/swagger/index.html")
	})
	port := fmt.Sprintf(":%s", utils.GetEnv("PORT", defaultPort))
	r.Run(port) // listen and serve on 0.0.0.0:8080
}
