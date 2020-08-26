package router

import (
	"gateway/controller"
	"gateway/lib"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"net/http"
)

func InitRouter(swag bool) *gin.Engine {

	router := gin.Default()
	//router.Use(middlewares...)
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	if swag {
		router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	router.StaticFS("/static", http.Dir(lib.GetDefaultConfBase().Base.StaticPath))

	Register(router, &controller.AdminLoginController{})

	Register(router, &controller.AdminController{})

	Register(router, &controller.ServiceController{})

	Register(router, &controller.PublicController{})

	Register(router, &controller.StatisticsController{})

	Register(router, &controller.AppController{})

	return router
}
