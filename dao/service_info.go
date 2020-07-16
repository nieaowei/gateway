package dao

import (
	"gateway/public"
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

func (p *ServiceInfo) ServiceDetail() string {
	return "service_info"
}

func (p *ServiceInfo) PageList(c *gin.Context, tx *gorm.DB, params *PageSize) (list []ServiceInfo, count uint, err error) {
	offset := (params.No - 1) * params.Size
	query := tx.SetCtx(public.GetTraceContext(c)).Model(p)
	if params.Info != "" {
		query = query.Where("service_name like ? or service_desc like ?", "%"+params.Info+"%", "%"+params.Info+"%")
	}

	err = query.Limit(params.Size).Offset(offset).Order("id desc").Find(&list).Error

	if err != nil {
		return
	}
	query.Count(&count)
	return
}

func (p *ServiceInfo) DeleteOne(c *gin.Context, tx *gorm.DB) (err error) {
	return tx.Delete(p).Error
}
