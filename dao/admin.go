package dao

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"time"
)

//type Admin struct {
//	gorm.Model
//	Salt     string `json:"salt"`
//	Username string `json:"username"`
//	Password string `json:"password"`
//}

type AdminSessionInfo struct {
	ID        uint      `json:"id"`
	Username  string    `json:"username"`
	LoginTime time.Time `json:"login_time"`
}

func (p *Admin) FindOne(c *gin.Context, tx *gorm.DB) (out *Admin, err error) {
	out = &Admin{}
	result := tx.Where(p).First(out)
	err = ErrorHandleForDB(result)
	return
}

func (p *Admin) UpdateAllByID(c *gin.Context, tx *gorm.DB) (err error) {
	res := tx.Omit("id", "created_at", "deleted_at").Save(p)
	err = ErrorHandleForDB(res)
	return
}

func (p *Admin) UpdateByID(c *gin.Context, tx *gorm.DB) (err error) {
	res := tx.Model(p).Omit("id", "created_at", "deleted_at").Updates(p)
	err = ErrorHandleForDB(res)
	return
}

func (p *Admin) DeleteByID(c *gin.Context, tx *gorm.DB) (err error) {
	result := tx.Delete(p)
	err = ErrorHandleForDB(result)
	return
}
