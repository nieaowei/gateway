package dao

import (
	"gateway/public"
	"github.com/e421083458/gorm"
	"github.com/gin-gonic/gin"
)

type ServiceLoadBalance struct {
	ID                     uint   `json:"id"`
	ServiceId              uint   `json:"service_id"`
	CheckMethod            uint   `json:"check_method"`
	CheckTimeout           uint   `json:"check_timeout"`
	CheckInterval          uint   `json:"check_interval"`
	RoundType              uint8  `json:"round_type"`
	IpList                 string `json:"ip_list"`
	WeightList             string `json:"weight_list"`
	ForbidLIst             string `json:"forbid_l_ist"`
	UpstreamConnectTimeout uint16 `json:"upstream_connect_timeout"`
	UpstreamHeaderTimeout  uint16 `json:"upstream_header_timeout"`
	UpstreamIdleTimeout    uint16 `json:"upstream_idle_timeout"`
	UpstreamMaxIdle        uint16 `json:"upstream_max_idle"`
}

func (p *ServiceLoadBalance) FindOne(c *gin.Context, tx *gorm.DB) (out *ServiceLoadBalance, err error) {
	out = &ServiceLoadBalance{}
	err = tx.SetCtx(public.GetTraceContext(c)).Where(p).First(out).Error
	if err != nil {
		return nil, err
	}
	return
}

func (p *ServiceLoadBalance) Save(c *gin.Context, tx *gorm.DB) (err error) {
	err = tx.Omit("id").SetCtx(public.GetTraceContext(c)).Save(p).Error
	if err != nil {
		return
	}
	return
}
