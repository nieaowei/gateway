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
//	MetadataTransform string `json:"header_transfor"`
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

func (p *ServiceHTTPRule) FindOneScan(c *gin.Context, tx *gorm.DB, out interface{}) (err error) {
	//out = &ServiceInfo{}
	result := tx.Model(p).Where(p).Limit(1).Scan(out)
	err = ErrorHandleForDB(result)
	return
}

func (p *ServiceHTTPRule) UpdateAllByServiceID(c *gin.Context, db *gorm.DB) (err error) {
	result := db.Select(GetFields(p)).Where("service_id=?", p.ServiceID).Updates(p)
	err = ErrorHandleForDB(result)
	return
}

func (p *ServiceHTTPRule) DeleteByID(c *gin.Context, tx *gorm.DB) (err error) {
	result := tx.Where(p).Delete(p)
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
		if err != gorm.ErrRecordNotFound {
			return errors.New("Integrity violation constraint")
		}
		ServiceTCPRule := &ServiceTCPRule{
			ServiceID: p.ServiceID,
		}
		err = db.First(ServiceTCPRule, ServiceTCPRule).Error
		if err != gorm.ErrRecordNotFound {
			return errors.New("Integrity violation constraint")
		}

		//check foreign
		serviceInfo := &ServiceInfo{
			ID: p.ServiceID,
		}
		err = db.First(serviceInfo, serviceInfo).Error
		if err != nil {
			return errors.New("In violation of the foreign key constraints")
		}
		// check unique ServiceName and suffix
		serviceHTTPRule := &ServiceHTTPRule{
			ServiceID: p.ServiceID,
		}
		orServiceHTTPRule := &ServiceHTTPRule{
			Rule: p.Rule,
		}
		res := db.Where(serviceHTTPRule).Or(orServiceHTTPRule).Limit(1).Find(serviceHTTPRule)
		if (res.Error != nil && res.Error != gorm.ErrRecordNotFound) || res.RowsAffected != 0 {
			return errors.New("Violation of the uniqueness constraint #Rule or #ServiceID")
		}
	}
	// make sure insert
	p.ID = 0
	err = db.Create(p).Error
	return err
}
