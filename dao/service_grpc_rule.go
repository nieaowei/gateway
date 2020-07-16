package dao

import (
	"gateway/public"
	"github.com/e421083458/gorm"
	"github.com/gin-gonic/gin"
)

type ServiceGrpcRule struct {
	ID             uint   `json:"id"`
	ServiceId      uint   `json:"service_id"`
	Port           uint16 `json:"port"`
	HeaderTransfor string `json:"header_transfor"`
}

func (p *ServiceGrpcRule) FindOne(c *gin.Context, tx *gorm.DB) (out *ServiceGrpcRule, err error) {
	out = &ServiceGrpcRule{}
	err = tx.SetCtx(public.GetTraceContext(c)).Where(p).First(out).Error
	if err != nil {
		return nil, err
	}
	return
}

func (p *ServiceGrpcRule) Save(c *gin.Context, tx *gorm.DB) (err error) {
	err = tx.Omit("id").SetCtx(public.GetTraceContext(c)).Save(p).Error
	if err != nil {
		return
	}
	return
}
