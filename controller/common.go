package controller

import "github.com/gin-gonic/gin"

type Controller interface {
	RouterRegister(group *gin.RouterGroup)
	RouterGroupName() (name string)
	Middlewares() (middlewares []gin.HandlerFunc)
}
