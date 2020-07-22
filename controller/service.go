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
	group.GET("/delete", service.ServiceDelete)
	group.POST("/add/http", service.ServiceAddHttp)
}

func (p *ServiceController) ServiceAddHttp(c *gin.Context) {
	params := &dto.ServiceAddHttpInput{}
	if err := params.BindValidParam(c); err != nil {
		//middleware.ResponseError(c, 1001, err)
		//return
	}
	//pass
	err := params.AddHttpService(c)
	if err != nil {
		middleware.ResponseError(c, 1002, err)
		return
	}
	middleware.ResponseSuccess(c, params)
	return
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

func (p *ServiceController) ServiceDelete(c *gin.Context) {
	params := &dto.ServiceDeleteInput{}
	if err := params.BindValidParam(c); err != nil {
		middleware.ResponseError(c, 1001, err)
		return
	}
	//pass
	err := params.Delete(c)
	if err != nil {
		middleware.ResponseError(c, 1002, err)
		return
	}
	middleware.ResponseSuccess(c, "")
	return
}
