package dao

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

//type ServiceHTTPRule struct {
//	gorm.Model
//	ServiceID      uint   `json:"service_id"`
//	RuleType       uint8  `json:"rule_type" validate:"oneof=0 1"`
//	Rule           string `json:"rule"`
//	NeedHttps      uint8  `json:"need_https" validate:"oneof=0 1"`
//	NeedStripUri   uint8  `json:"need_strip_uri" validate:"oneof=0 1"`
//	NeedWebsocket  uint8  `json:"need_websocket" validate:"oneof=0 1"`
//	UrlRewrite     string `json:"url_rewrite" validate:"valid_url_rewrite"`
//	HeaderTransfor string `json:"header_transfor"`
//}

func (p *ServiceHTTPRule) BeforeUpdate(tx *gorm.DB) error {
	tx = tx.Statement.Where("deleted_at IS NULL").Omit("created_at", "service_id", "deleted_at", "id")
	return nil
}

func (p *ServiceHTTPRule) BeforeDelete(tx *gorm.DB) error {
	tx = tx.Statement.Where("deleted_at IS NULL")
	return nil
}

func (p *ServiceHTTPRule) BeforeCreate(db *gorm.DB) error {
	db.Statement.Omit("id")
	return nil
}

func (p *ServiceHTTPRule) FindOne(c *gin.Context, tx *gorm.DB) (out *ServiceHTTPRule, err error) {
	out = &ServiceHTTPRule{}
	result := tx.Where(p).First(out)
	err = ErrorHandleForDB(result)
	return
}

func (p *ServiceHTTPRule) UpdateAllByServiceID(c *gin.Context, db *gorm.DB) (err error) {
	result := db.Select(GetFields(p)).Where("service_id=?", p.ServiceID).Updates(p)
	err = ErrorHandleForDB(result)
	return
}

func (p *ServiceHTTPRule) DeleteByID(c *gin.Context, tx *gorm.DB) (err error) {
	result := tx.Delete(p)
	err = ErrorHandleForDB(result)
	return
}

func (p *ServiceHTTPRule) InsertAfterCheck(c *gin.Context, db *gorm.DB, check bool) (err error) {
	if check {
		//check integrity
		serviceGrpcRule := &ServiceGrpcRule{
			ServiceID: p.ServiceID,
		}
		err = db.First(serviceGrpcRule, serviceGrpcRule).Error
		if err != nil && err != gorm.ErrRecordNotFound {
			return errors.New("Integrity violation constraint")
		}
		ServiceTCPRule := &ServiceTCPRule{
			ServiceID: p.ServiceID,
		}
		err = db.First(ServiceTCPRule, ServiceTCPRule).Error
		if err != nil && err != gorm.ErrRecordNotFound {
			return errors.New("Integrity violation constraint")
		}

		//check foregin
		serviceInfo := &ServiceInfo{
			Model: gorm.Model{ID: p.ServiceID},
		}
		err = db.First(serviceInfo, serviceInfo).Error
		if err != nil && err != gorm.ErrRecordNotFound {
			return errors.New("In violation of the foreign key constraints")
		}
		ServiceHTTPRule := &ServiceHTTPRule{
			ServiceID: p.ServiceID,
		}
		// check unique ServiceName
		err = db.First(ServiceHTTPRule, ServiceHTTPRule).Error
		if err != nil && err != gorm.ErrRecordNotFound {
			return errors.New("Violation of the uniqueness constraint")
		}
	}
	// make sure insert
	p.ID = 0
	err = db.Create(p).Error
	return err
}
