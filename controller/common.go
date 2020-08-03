package controller

import "github.com/gin-gonic/gin"

type Controller interface {
	RouterRegister(group *gin.RouterGroup)
	RouterGroupName() string
	Middleware() []gin.HandlerFunc
}
