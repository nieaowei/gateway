package dao

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

//type ServiceTCPRule struct {
//	gorm.Model
//	ServiceID uint   `json:"service_id"`
//	Port      uint16 `json:"port"`
//}

func (p *ServiceTCPRule) BeforeUpdate(tx *gorm.DB) error {
	tx = tx.Statement.Where("deleted_at IS NULL").Omit("created_at", "service_id", "deleted_at", "id")
	return nil
}

func (p *ServiceTCPRule) BeforeDelete(tx *gorm.DB) error {
	tx = tx.Statement.Where("deleted_at IS NULL")
	return nil
}

func (p *ServiceTCPRule) BeforeCreate(db *gorm.DB) error {
	db.Statement.Omit("id")
	return nil
}

func (p *ServiceTCPRule) FindOne(c *gin.Context, tx *gorm.DB) (out *ServiceTCPRule, err error) {
	out = &ServiceTCPRule{}
	result := tx.Where(p).First(out)
	err = ErrorHandleForDB(result)
	return
}

func (p *ServiceTCPRule) FindOneScan(c *gin.Context, tx *gorm.DB, out interface{}) (err error) {
	//out = &ServiceInfo{}
	result := tx.Model(p).Where(p).Limit(1).Scan(out)
	err = ErrorHandleForDB(result)
	return
}

func (p *ServiceTCPRule) UpdateAllByServiceID(c *gin.Context, db *gorm.DB) (err error) {
	result := db.Select(GetFields(p)).Where("service_id=?", p.ServiceID).Updates(p)
	err = ErrorHandleForDB(result)
	return
}

func (p *ServiceTCPRule) DeleteByID(c *gin.Context, tx *gorm.DB) (err error) {
	result := tx.Delete(p)
	err = ErrorHandleForDB(result)
	return
}

func (p *ServiceTCPRule) InsertAfterCheck(c *gin.Context, db *gorm.DB, check bool) (err error) {
	if check {
		//check integrity
		ServiceHTTPRule := &ServiceHTTPRule{
			ServiceID: p.ServiceID,
		}
		err = db.First(ServiceHTTPRule, ServiceHTTPRule).Error
		if err != nil && err != gorm.ErrRecordNotFound {
			return errors.New("Integrity violation constraint")
		}

		serviceGrpcRule := &ServiceGrpcRule{
			ServiceID: p.ServiceID,
		}
		err = db.First(serviceGrpcRule, serviceGrpcRule).Error
		if err != nil && err != gorm.ErrRecordNotFound {
			return errors.New("Integrity violation constraint")
		}

		//check foregin
		serviceInfo := &ServiceInfo{
			ID: p.ServiceID,
		}
		err = db.First(serviceInfo, serviceInfo).Error
		if err != nil && err != gorm.ErrRecordNotFound {
			return errors.New("In violation of the foreign key constraints")
		}

		// check unique ServiceID
		ServiceTCPRule := &ServiceTCPRule{
			ServiceID: p.ServiceID,
		}
		err = db.First(ServiceTCPRule, ServiceTCPRule).Error
		if err != nil && err != gorm.ErrRecordNotFound {
			return errors.New("Violation of the uniqueness constraint")
		}
	}
	// make sure insert
	p.ID = 0
	err = db.Create(p).Error
	return err
}
