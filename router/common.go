package router

import (
	"gateway/controller"
	"github.com/gin-gonic/gin"
)

func Register(router *gin.Engine, c controller.Controller) {
	group := router.Group(c.RouterGroupName(), c.Middleware()...)
	c.RouterRegister(group)
}
