package dao

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
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

func (p *App) PageListIdDesc(c *gin.Context, tx *gorm.DB, params *PageSize) (list []App, count int64, err error) {
	offset := (params.No - 1) * params.Size
	query := tx
	if params.Info != "" {
		query = query.Where("name like ?", "%"+params.Info+"%")
	}
	query.Model(p).Count(&count)

	err = query.Limit(params.Size).Offset(offset).Order("id desc").Find(&list).Error

	if err != nil {
		return
	}
	return
}

func (p *App) InsertAfterCheck(c *gin.Context, db *gorm.DB, check bool) (err error) {
	if check {
		app := &App{
			AppID: p.AppID,
		}
		// check unique AppID
		err = db.First(app, app).Error
		if err != gorm.ErrRecordNotFound {
			return errors.New("Violation of the uniqueness constraint #App.AppID")
		}
	}
	// make sure insert
	res := db.Create(p)
	err = ErrorHandleForDB(res)
	return
}
