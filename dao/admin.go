package dao

import (
	"gateway/public"
	"github.com/e421083458/gorm"
	"github.com/gin-gonic/gin"
	"time"
)

type Admin struct {
	gorm.Model
	Salt     string `json:"salt"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type AdminSessionInfo struct {
	ID        uint      `json:"id"`
	Username  string    `json:"username"`
	LoginTime time.Time `json:"login_time"`
}

func (p *Admin) FindOne(c *gin.Context, tx *gorm.DB) (out *Admin, err error) {
	out = &Admin{}
	err = tx.SetCtx(public.GetTraceContext(c)).Where(p).First(out).Error
	if err != nil {
		return nil, err
	}
	return
}

func (p *Admin) Save(c *gin.Context, tx *gorm.DB) (err error) {
	err = tx.Omit("id", "created_at").SetCtx(public.GetTraceContext(c)).Save(p).Error
	if err != nil {
		return
	}
	return
}

func (p *Admin) Updates(c *gin.Context, tx *gorm.DB) (err error) {
	err = tx.Model(p).Omit("id").SetCtx(public.GetTraceContext(c)).Updates(p).Error
	if err != nil {
		return
	}
	return
}
