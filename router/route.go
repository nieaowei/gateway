package router

import (
	"gateway/controller"
	"github.com/gin-gonic/gin"
)

func InitRouter(middlewares ...gin.HandlerFunc) *gin.Engine {

	router := gin.Default()
	router.Use(middlewares...)
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	Register(router, &controller.AdminLoginController{})

	Register(router, &controller.AdminController{})

	Register(router, &controller.ServiceController{})

	return router
}
