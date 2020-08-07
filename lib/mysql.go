package lib

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

var DefaultDB *gorm.DB

func InitDBPool() {
	//普通的db方式G
	var err error
	conf := GetDefaultConfMysql()
	DefaultDB, err = gorm.Open(mysql.Open(conf.DataSourceName), &gorm.Config{
		Logger: logger.New(log.New(os.Stdout, "\r\n", log.LstdFlags), logger.Config{
			SlowThreshold: 100 * time.Millisecond,
			LogLevel:      logger.Info,
			Colorful:      true,
		}),
	})
	if err != nil {
		panic(err)
	}
}
func GetDefaultDB() (db *gorm.DB) {
	if DefaultDB == nil {
		//panic("db not init")
		InitDBPool()
	}
	return DefaultDB
}

func CloseDB() {

}
