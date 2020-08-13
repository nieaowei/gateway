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

func (p *ServiceAccessControl) BeforeUpdate(tx *gorm.DB) error {
	tx = tx.Statement.Where("deleted_at IS NULL").Omit("created_at", "service_id", "deleted_at", "id")
	return nil
}

func (p *ServiceAccessControl) BeforeDelete(tx *gorm.DB) error {
	tx = tx.Statement.Where("deleted_at IS NULL")
	return nil
}

func (p *ServiceAccessControl) BeforeCreate(db *gorm.DB) error {
	db.Statement.Omit("id")
	return nil
}

func (p *ServiceAccessControl) FindOne(c *gin.Context, tx *gorm.DB) (out *ServiceAccessControl, err error) {
	out = &ServiceAccessControl{}
	result := tx.Where(p).First(out)
	err = ErrorHandleForDB(result)
	return
}

func (p *ServiceAccessControl) FindOneScan(c *gin.Context, tx *gorm.DB, out interface{}) (err error) {
	//out = &ServiceInfo{}
	result := tx.Model(p).Where(p).Limit(1).Scan(out)
	err = ErrorHandleForDB(result)
	return
}

func (p *ServiceAccessControl) UpdateAllByServiceID(c *gin.Context, db *gorm.DB) (err error) {
	//data := map[string]interface{}{}
	//data["ServiceID"] = p.ServiceID
	//data["OpenAuth"] = p.OpenAuth
	//data["BlackList"] = p.BlackList
	//data["WhiteList"] = p.WhiteList
	//data["WhiteHostName"] = p.WhiteHostName
	//data["ClientipFlowLimit"] = p.ClientipFlowLimit
	//data["ServiceFlowLimit"] = p.ServiceFlowLimit
	//return nil

	result := db.Select(GetFields(p)).Where("service_id=?", p.ServiceID).Updates(p)
	err = ErrorHandleForDB(result)
	return
}

func (p *ServiceAccessControl) DeleteByID(c *gin.Context, tx *gorm.DB) (err error) {
	result := tx.Where(p).Delete(p)
	err = ErrorHandleForDB(result)
	return
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
	result := tx.Create(p)
	err = ErrorHandleForDB(result)
	return
}
