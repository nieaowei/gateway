package dao

import (
	"gateway/public"
	"github.com/e421083458/gorm"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

type ServiceHttpRule struct {
	gorm.Model
	ServiceId      uint   `json:"service_id"`
	RuleType       uint8  `json:"rule_type" validate:"oneof=0 1"`
	Rule           string `json:"rule"`
	NeedHttps      uint8  `json:"need_https" validate:"oneof=0 1"`
	NeedStripUri   uint8  `json:"need_strip_uri" validate:"oneof=0 1"`
	NeedWebsocket  uint8  `json:"need_websocket" validate:"oneof=0 1"`
	UrlRewrite     string `json:"url_rewrite" validate:"valid_url_rewrite"`
	HeaderTransfor string `json:"header_transfor"`
}

func (p *ServiceHttpRule) FindOne(c *gin.Context, tx *gorm.DB) (out *ServiceHttpRule, err error) {
	out = &ServiceHttpRule{}
	err = tx.SetCtx(public.GetTraceContext(c)).Where(p).First(out).Error
	if err != nil {
		return nil, err
	}
	return
}

func (p *ServiceHttpRule) Save(c *gin.Context, tx *gorm.DB) (err error) {
	err = tx.Omit("id").SetCtx(public.GetTraceContext(c)).Save(p).Error
	if err != nil {
		return
	}
	return
}

func (p *ServiceHttpRule) Delete(c *gin.Context, tx *gorm.DB) (err error) {

	return tx.Where(p).Delete(p).Error
}

func (p *ServiceHttpRule) InsertAfterCheck(c *gin.Context, db *gorm.DB, check bool) (err error) {
	if check {
		//check integrity
		serviceGrpcRule := &ServiceGrpcRule{
			ServiceId: p.ServiceId,
		}
		err = db.First(serviceGrpcRule, serviceGrpcRule).Error
		if err != nil && err != gorm.ErrRecordNotFound {
			return errors.New("Integrity violation constraint")
		}
		serviceTcpRule := &ServiceTcpRule{
			ServiceId: p.ServiceId,
		}
		err = db.First(serviceTcpRule, serviceTcpRule).Error
		if err != nil && err != gorm.ErrRecordNotFound {
			return errors.New("Integrity violation constraint")
		}

		//check foregin
		serviceInfo := &ServiceInfo{
			Model: gorm.Model{ID: p.ServiceId},
		}
		err = db.First(serviceInfo, serviceInfo).Error
		if err != nil && err != gorm.ErrRecordNotFound {
			return errors.New("In violation of the foreign key constraints")
		}
		serviceHttpRule := &ServiceHttpRule{
			ServiceId: p.ServiceId,
		}
		// check unique ServiceName
		err = db.First(serviceHttpRule, serviceHttpRule).Error
		if err != nil && err != gorm.ErrRecordNotFound {
			return errors.New("Violation of the uniqueness constraint")
		}
	}
	// make sure insert
	p.ID = 0
	err = db.Create(p).Error
	return err
}
