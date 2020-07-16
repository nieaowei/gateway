package dao

import (
	"gateway/public"
	"github.com/e421083458/gorm"
	"github.com/gin-gonic/gin"
)

type ServiceTcpRule struct {
	ID        uint   `json:"id"`
	ServiceId uint   `json:"service_id"`
	Port      uint16 `json:"port"`
}

func (p *ServiceTcpRule) FindOne(c *gin.Context, tx *gorm.DB) (out *ServiceTcpRule, err error) {
	out = &ServiceTcpRule{}
	err = tx.SetCtx(public.GetTraceContext(c)).Where(p).First(out).Error
	if err != nil {
		return nil, err
	}
	return
}

func (p *ServiceTcpRule) Save(c *gin.Context, tx *gorm.DB) (err error) {
	err = tx.Omit("id").SetCtx(public.GetTraceContext(c)).Save(p).Error
	if err != nil {
		return
	}
	return
}

func (p *ServiceTcpRule) Delete(c *gin.Context, tx *gorm.DB) (err error) {

	return tx.Delete(p).Error
}
