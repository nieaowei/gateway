package router

import (
	"gateway/controller"
	"gateway/lib"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRouter(middlewares ...gin.HandlerFunc) *gin.Engine {

	router := gin.Default()
	router.Use(middlewares...)
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	router.StaticFS("/static", http.Dir(lib.GetDefaultConfBase().Base.StaticPath))

	Register(router, &controller.AdminLoginController{})

	Register(router, &controller.AdminController{})

	Register(router, &controller.ServiceController{})

	Register(router, &controller.PublicController{})

	return router
}
