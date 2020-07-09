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
	err = tx.SetCtx(public.GetTraceContext(c)).Where(p).Find(out).Error
	if err != nil {
		return nil, err
	}
	return
}
