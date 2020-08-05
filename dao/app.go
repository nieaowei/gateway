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
	result := tx.Where(p).First(out)
	err = ErrorHandleForDB(result)
	return
}

func (p *App) FindOneScan(c *gin.Context, tx *gorm.DB, out interface{}) (err error) {
	//out = &ServiceInfo{}
	result := tx.Model(p).Where(p).Limit(1).Scan(out)
	err = ErrorHandleForDB(result)
	return
}

func (p *App) Save(c *gin.Context, tx *gorm.DB) (err error) {
	res := tx.Omit("id", "created_at").Save(p)
	err = ErrorHandleForDB(res)
	return
}

func (p *App) DeleteByID(c *gin.Context, tx *gorm.DB) (err error) {
	result := tx.Delete(p)
	err = ErrorHandleForDB(result)
	return
}
