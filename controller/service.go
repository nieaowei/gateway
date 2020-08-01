package controller

import (
	"gateway/dto"
	"github.com/gin-gonic/gin"
)

type ServiceController struct {
}

func ServiceRigster(group *gin.RouterGroup) {
	service := &ServiceController{}
	group.GET("/list", service.ServiceList)
	group.GET("/del", service.DeleteService)
	group.POST("/http/add", service.AddHttpService)
	group.POST("/http/update", service.UpdateHttpService)
}

func (p *ServiceController) AddHttpService(c *gin.Context) {
	params := &dto.AddHttpServiceInput{}
	if err := params.BindValidParam(c); err != nil {
		ResponseError(c, 1001, err)
		return
	}
	//pass
	err := params.AddHttpService(c)
	if err != nil {
		ResponseError(c, 1002, err)
		return
	}
	ResponseSuccess(c, params)
	return
}

func (p *ServiceController) UpdateHttpService(c *gin.Context) {
	params := &dto.UpdateHttpServiceInput{}
	if err := params.BindValidParam(c); err != nil {
		ResponseError(c, 1001, err)
		return
	}
	//pass
	err := params.UpdateHttpService(c)
	if err != nil {
		ResponseError(c, 1002, err)
		return
	}
	ResponseSuccess(c, params)
	return
}

func (p *ServiceController) ServiceList(c *gin.Context) {
	params := &dto.ServiceListInput{}
	if err := params.BindValidParam(c); err != nil {
		ResponseError(c, 1001, err)
		return
	}
	//pass
	out, err := params.GetServiceList(c)
	if err != nil {
		ResponseError(c, 1002, err)
		return
	}
	ResponseSuccess(c, out)
	return
}

func (p *ServiceController) DeleteService(c *gin.Context) {
	params := &dto.DeleteServiceInput{}
	if err := params.BindValidParam(c); err != nil {
		ResponseError(c, 1001, err)
		return
	}
	//pass
	err := params.Delete(c)
	if err != nil {
		ResponseError(c, 1002, err)
		return
	}
	ResponseSuccess(c, "")
	return
}
