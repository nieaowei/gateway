package router

import (
	"errors"
	"gateway/controller"
	"gateway/dto"
	"gateway/lib"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"net/http"
)

func InitRouter(swag bool) *gin.Engine {

	router := gin.Default()
	router.LoadHTMLGlob("./templates/*")
	//router.Use(middlewares...)
	router.NoRoute(func(c *gin.Context) {
		dto.ResponseError(c, 404, errors.New("Not found"))
	})
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "welcome.html", gin.H{
			"title":        "微服务仪表盘API接口",
			"service_name": "Dashboard Service",
			"welcome_msg":  "The service is started.",
			"items": []gin.H{
				{
					"name": "API Document",
					"link": "swagger/index.html",
					"tag":  "Swagger",
				},
			},
		})
	})
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
