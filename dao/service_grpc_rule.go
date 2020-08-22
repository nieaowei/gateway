package dao

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

//type ServiceGrpcRule struct {
//	gorm.Model
//	ServiceID      uint   `json:"service_id"`
//	Port           uint16 `json:"port"`
//	MetadataTransform string `json:"header_transfor"`
//}
func (p *ServiceGrpcRule) BeforeUpdate(tx *gorm.DB) error {
	tx = tx.Statement.Where("deleted_at IS NULL").Omit("created_at", "service_id", "deleted_at", "id")
	return nil
}

func (p *ServiceGrpcRule) BeforeDelete(tx *gorm.DB) error {
	tx = tx.Statement.Where("deleted_at IS NULL")
	return nil
}

func (p *ServiceGrpcRule) BeforeCreate(db *gorm.DB) error {
	db.Statement.Omit("id")
	return nil
}

func (p *ServiceGrpcRule) FindOne(c *gin.Context, tx *gorm.DB) (out *ServiceGrpcRule, err error) {
	out = &ServiceGrpcRule{}
	result := tx.Where(p).First(out)
	err = ErrorHandleForDB(result)
	return
}

func (p *ServiceGrpcRule) FindOneScan(c *gin.Context, tx *gorm.DB, out interface{}) (err error) {
	//out = &ServiceInfo{}
	result := tx.Model(p).Where(p).Limit(1).Scan(out)
	err = ErrorHandleForDB(result)
	return
}

func (p *ServiceGrpcRule) UpdateAllByServiceID(c *gin.Context, db *gorm.DB) (err error) {
	result := db.Select(GetFields(p)).Where("service_id=?", p.ServiceID).Updates(p)
	err = ErrorHandleForDB(result)
	return
}

func (p *ServiceGrpcRule) DeleteByID(c *gin.Context, tx *gorm.DB) (err error) {
	result := tx.Where(p).Delete(p)
	err = ErrorHandleForDB(result)
	return
}

func (p *ServiceGrpcRule) InsertAfterCheck(c *gin.Context, db *gorm.DB, check bool) (err error) {
	if check {
		//check integrity
		ServiceHTTPRule := &ServiceHTTPRule{
			ServiceID: p.ServiceID,
		}
		err = db.First(ServiceHTTPRule, ServiceHTTPRule).Error
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
			ID: p.ServiceID,
		}
		err = db.First(serviceInfo, serviceInfo).Error
		if err != nil && err != gorm.ErrRecordNotFound {
			return errors.New("In violation of the foreign key constraints")
		}

		// check unique ServiceID
		serviceGrpcRule := &ServiceGrpcRule{
			ServiceID: p.ServiceID,
		}
		err = db.First(serviceGrpcRule, serviceGrpcRule).Error
		if err != nil && err != gorm.ErrRecordNotFound {
			return errors.New("Violation of the uniqueness constraint")
		}
	}
	// make sure insert
	p.ID = 0
	err = db.Create(p).Error
	return err
}
