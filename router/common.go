package router

import (
	"github.com/gin-gonic/gin"
)

type Controller interface {
	Register(group *gin.RouterGroup)
	GroupName() string
	Middleware() []gin.HandlerFunc
}

func Register(router *gin.Engine, c Controller) {
	group := router.Group(c.GroupName(), c.Middleware()...)
	c.Register(group)
}
