package dao

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

//type App struct {
//	gorm.Model
//	AppId    string `json:"app_id"`
//	Name     string `json:"name"`
//	Secret   string `json:"secret"`
//	WhiteIps string `json:"white_ips"`
//	qpd      uint   `json:"qpd"`
//	qps      uint   `json:"qps"`
//}

func (p *App) FindOne(c *gin.Context, tx *gorm.DB) (out *App, err error) {
	out = &App{}
	err = tx.Where(p).First(out).Error
	if err != nil {
		return nil, err
	}
	return
}

func (p *App) Save(c *gin.Context, tx *gorm.DB) (err error) {
	err = tx.Omit("id", "created_at").Save(p).Error
	if err != nil {
		return
	}
	return
}

func (p *App) Delete(c *gin.Context, tx *gorm.DB) (err error) {
	return tx.Delete(p).Error
}
