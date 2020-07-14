package controller

import (
	"gateway/dto"
	"gateway/middleware"
	"github.com/gin-gonic/gin"
)

type ServiceController struct {
}

func ServiceRigster(group *gin.RouterGroup) {
	service := &ServiceController{}
	group.GET("/list", service.ServiceList)
}

func (p *ServiceController) ServiceList(c *gin.Context) {
	params := &dto.ServiceListInput{}
	if err := params.BindValidParam(c); err != nil {
		middleware.ResponseError(c, 1001, err)
		return
	}
	//pass
	out, err := params.GetServiceList(c)
	if err != nil {
		middleware.ResponseError(c, 1002, err)
		return
	}
	middleware.ResponseSuccess(c, out)
	return
}
