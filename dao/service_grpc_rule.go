package dao

import (
	"gateway/public"
	"github.com/e421083458/gorm"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

type ServiceGrpcRule struct {
	gorm.Model
	ServiceId      uint   `json:"service_id"`
	Port           uint16 `json:"port"`
	HeaderTransfor string `json:"header_transfor"`
}

func (p *ServiceGrpcRule) FindOne(c *gin.Context, tx *gorm.DB) (out *ServiceGrpcRule, err error) {
	out = &ServiceGrpcRule{}
	err = tx.SetCtx(public.GetTraceContext(c)).Where(p).First(out).Error
	if err != nil {
		return nil, err
	}
	return
}

func (p *ServiceGrpcRule) Save(c *gin.Context, tx *gorm.DB) (err error) {
	err = tx.Omit("id").SetCtx(public.GetTraceContext(c)).Save(p).Error
	if err != nil {
		return
	}
	return
}

func (p *ServiceGrpcRule) Delete(c *gin.Context, tx *gorm.DB) (err error) {
	return tx.Where(p).Delete(p).Error
}

func (p *ServiceGrpcRule) AddAfterCheck(c *gin.Context, db *gorm.DB, check bool) (err error) {
	if check {
		//check integrity
		serviceHttpRule := &ServiceHttpRule{
			ServiceId: p.ServiceId,
		}
		err = db.First(serviceHttpRule, serviceHttpRule).Error
		if err != gorm.ErrRecordNotFound {
			return errors.New("Integrity violation constraint")
		}

		serviceTcpRule := &ServiceTcpRule{
			ServiceId: p.ServiceId,
		}
		err = db.First(serviceTcpRule, serviceTcpRule).Error
		if err != nil {
			return errors.New("Integrity violation constraint")
		}

		//check foregin
		serviceInfo := &ServiceInfo{
			Model: gorm.Model{ID: p.ServiceId},
		}
		err = db.First(serviceInfo, serviceInfo).Error
		if err != nil {
			return errors.New("In violation of the foreign key constraints")
		}

		// check unique ServiceId
		serviceGrpcRule := &ServiceGrpcRule{
			ServiceId: p.ServiceId,
		}
		err = db.First(serviceGrpcRule, serviceGrpcRule).Error
		if err != nil {
			return errors.New("Violation of the uniqueness constraint")
		}
	}
	// make sure insert
	p.ID = 0
	err = db.Create(p).Error
	return err
}
