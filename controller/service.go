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
	group.GET("/detail", service.GetServiceDetail)

}

func (p *ServiceController) GetServiceDetail(c *gin.Context) {
	dto.Exec(&dto.GetServiceDetailInput{}, c)
	return
}

func (p *ServiceController) AddHttpService(c *gin.Context) {
	dto.Exec(&dto.AddHttpServiceInput{}, c)
	return
}

func (p *ServiceController) UpdateHttpService(c *gin.Context) {
	dto.Exec(&dto.UpdateHttpServiceInput{}, c)
	return
}

func (p *ServiceController) ServiceList(c *gin.Context) {
	dto.Exec(&dto.ServiceListInput{}, c)
	return
}

func (p *ServiceController) DeleteService(c *gin.Context) {
	dto.Exec(&dto.DeleteServiceInput{}, c)
	return
}
