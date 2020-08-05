package dao

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	. "strings"
)

//type ServiceLoadBalance struct {
//	gorm.Model
//	ServiceID              uint   `json:"service_id"`
//	CheckMethod            uint   `json:"check_method"`
//	CheckTimeout           uint   `json:"check_timeout"`
//	CheckInterval          uint   `json:"check_interval"`
//	RoundType              uint8  `json:"round_type" validate:"oneof=0 1 2"`
//	IpList                 string `json:"ip_list" validate:"valid_ip_list"`
//	WeightList             string `json:"weight_list" validate:"valid_weight_list"`
//	ForbidList             string `json:"forbid_list"`
//	UpstreamConnectTimeout uint16 `json:"upstream_connect_timeout"`
//	UpstreamHeaderTimeout  uint16 `json:"upstream_header_timeout"`
//	UpstreamIdleTimeout    uint16 `json:"upstream_idle_timeout"`
//	UpstreamMaxIdle        uint16 `json:"upstream_max_idle"`
//}

func (p *ServiceLoadBalance) BeforeUpdate(tx *gorm.DB) error {
	tx = tx.Statement.Where("deleted_at IS NULL").Omit("created_at", "service_id", "deleted_at", "id")
	return nil
}

func (p *ServiceLoadBalance) BeforeDelete(tx *gorm.DB) error {
	tx = tx.Statement.Where("deleted_at IS NULL")
	return nil
}

func (p *ServiceLoadBalance) BeforeCreate(db *gorm.DB) error {
	db.Statement.Omit("id")
	return nil
}

func (p *ServiceLoadBalance) FindOne(c *gin.Context, tx *gorm.DB) (out *ServiceLoadBalance, err error) {
	out = &ServiceLoadBalance{}
	result := tx.Where(p).First(out)
	err = ErrorHandleForDB(result)
	return
}

func (p *ServiceLoadBalance) FindOneScan(c *gin.Context, db *gorm.DB, out interface{}) (err error) {
	result := db.Model(p).Where(p).Limit(1).Scan(out)
	err = ErrorHandleForDB(result)
	return
}

func (p *ServiceLoadBalance) GetIPListByModel() (list []string) {
	return Split(p.IPList, ",")
}

func (p *ServiceLoadBalanceExceptModel) GetIPListByModel() (list []string) {
	return Split(p.IPList, ",")
}

func (p *ServiceLoadBalance) UpdateAllByServiceID(c *gin.Context, db *gorm.DB) (err error) {
	result := db.Select(GetFields(p)).Where("service_id=?", p.ServiceID).Updates(p)
	err = ErrorHandleForDB(result)
	return
}

func (p *ServiceLoadBalance) DeleteByID(c *gin.Context, tx *gorm.DB) (err error) {
	result := tx.Delete(p)
	err = ErrorHandleForDB(result)
	return
}

func (p *ServiceLoadBalance) InsertAfterCheck(c *gin.Context, tx *gorm.DB, check bool) (err error) {
	if check {
		// check unique ServiceID
		slb := &ServiceLoadBalance{
			ServiceID: p.ServiceID,
		}
		_, err = slb.FindOne(c, tx)
		if err != nil && err != gorm.ErrRecordNotFound {
			return
		}
		err = nil
	}
	return tx.Create(p).Error
}
