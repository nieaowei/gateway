package dao

import (
	"gateway/public"
	"github.com/e421083458/gorm"
	"github.com/gin-gonic/gin"
)

type ServiceAccessControl struct {
	gorm.Model
	ServiceId         uint   `json:"service_id"`
	OpenAuth          uint8  `json:"open_auth"`
	BlackList         string `json:"black_list"`
	WhiteList         string `json:"white_list"`
	WhiteHostName     string `json:"white_host_name"`
	ClientipFlowLimit uint16 `json:"clientip_flow_limit"`
	ServiceFlowLimit  uint16 `json:"service_flow_limit"`
}

func (p *ServiceAccessControl) FindOne(c *gin.Context, tx *gorm.DB) (out *ServiceAccessControl, err error) {
	out = &ServiceAccessControl{}
	err = tx.SetCtx(public.GetTraceContext(c)).Where(p).First(out).Error
	if err != nil {
		return nil, err
	}
	return
}

func (p *ServiceAccessControl) Save(c *gin.Context, tx *gorm.DB) (err error) {
	err = tx.Omit("id").SetCtx(public.GetTraceContext(c)).Save(p).Error
	if err != nil {
		return
	}
	return
}

func (p *ServiceAccessControl) Delete(c *gin.Context, tx *gorm.DB) (err error) {
	return tx.Delete(p).Error
}
