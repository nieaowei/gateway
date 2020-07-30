package dao

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

//type ServiceAccessControl struct {
//	gorm.Model
//	ServiceID         uint   `json:"service_id"`
//	OpenAuth          uint8  `json:"open_auth" validate:"oneof=0 1"`
//	BlackList         string `json:"black_list"`
//	WhiteList         string `json:"white_list"`
//	WhiteHostName     string `json:"white_host_name"`
//	ClientipFlowLimit uint16 `json:"clientip_flow_limit"`
//	ServiceFlowLimit  uint16 `json:"service_flow_limit"`
//}

func (p *ServiceAccessControl) FindOne(c *gin.Context, tx *gorm.DB) (out *ServiceAccessControl, err error) {
	out = &ServiceAccessControl{}
	err = tx.Where(p).First(out).Error
	if err != nil {
		return nil, err
	}
	return
}

func (p *ServiceAccessControl) Save(c *gin.Context, tx *gorm.DB) (err error) {
	err = tx.Omit("id").Save(p).Error
	if err != nil {
		return
	}
	return
}

func (p *ServiceAccessControl) Delete(c *gin.Context, tx *gorm.DB) (err error) {
	return tx.Where(p).Delete(p).Error
}

func (p *ServiceAccessControl) InsertAfterCheck(c *gin.Context, tx *gorm.DB, check bool) (err error) {
	if check {
		// check unique ServiceID
		asc := &ServiceAccessControl{
			ServiceID: p.ServiceID,
		}
		_, err = asc.FindOne(c, tx)
		if err != nil && err != gorm.ErrRecordNotFound {
			return
		}
		err = nil
	}
	return tx.Create(p).Error
}
