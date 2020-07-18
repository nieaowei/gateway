package dao

import (
	"gateway/public"
	"github.com/e421083458/gorm"
	"github.com/gin-gonic/gin"
)

type ServiceHttpRule struct {
	gorm.Model
	ServiceId      uint   `json:"service_id"`
	RuleType       uint8  `json:"rule_type"`
	Rule           string `json:"rule"`
	NeedHttps      uint8  `json:"need_https"`
	NeedStripUri   uint8  `json:"need_strip_uri"`
	NeedWebSocket  uint8  `json:"need_web_socket"`
	UrlRewrite     string `json:"url_rewrite"`
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

	return tx.Delete(p).Error
}
