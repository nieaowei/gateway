package dao

import (
	"gateway/public"
	"github.com/e421083458/gorm"
	"github.com/gin-gonic/gin"
	. "strings"
)

type ServiceLoadBalance struct {
	gorm.Model
	ServiceId              uint   `json:"service_id"`
	CheckMethod            uint   `json:"check_method"`
	CheckTimeout           uint   `json:"check_timeout"`
	CheckInterval          uint   `json:"check_interval"`
	RoundType              uint8  `json:"round_type" validate:"oneof=0 1 2"`
	IpList                 string `json:"ip_list" validate:"valid_ip_list"`
	WeightList             string `json:"weight_list" validate:"valid_weight_list"`
	ForbidList             string `json:"forbid_list"`
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

func (p *ServiceLoadBalance) GetIPListByModel() (list []string) {
	return Split(p.IpList, ",")
}

func (p *ServiceLoadBalance) Delete(c *gin.Context, tx *gorm.DB) (err error) {

	return tx.Where(p).Delete(p).Error
}

func (p *ServiceLoadBalance) InsertAfterCheck(c *gin.Context, tx *gorm.DB, check bool) (err error) {
	if check {
		// check unique ServiceId
		slb := &ServiceLoadBalance{
			ServiceId: p.ServiceId,
		}
		_, err = slb.FindOne(c, tx)
		if err != nil && err != gorm.ErrRecordNotFound {
			return
		}
		err = nil
	}
	return tx.Create(p).Error
}
