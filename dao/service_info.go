package dao

import (
	"gateway/dto"
	"github.com/e421083458/gorm"
	"github.com/gin-gonic/gin"
)

type ServiceInfo struct {
	gorm.Model
	LoadType    uint   `json:"load_type"`
	ServiceName string `json:"service_name"`
	ServiceDesc string `json:"service_desc"`
}

func (p *ServiceInfo) TableName() string {
	return "service_info"
}

func (p *ServiceInfo) PageList(c *gin.Context, tx *gorm.DB, params *dto.ServiceListInput) (list []ServiceInfo, count uint, err error) {

	return
}
