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

func InitDBPool() error {
	//普通的db方式
	var err error
	DefaultDB, err = gorm.Open(mysql.Open(GetDefaultConfMysql().DataSourceName), &gorm.Config{
		Logger: logger.New(log.New(os.Stdout, "\r\n", log.LstdFlags), logger.Config{
			SlowThreshold: 100 * time.Millisecond,
			LogLevel:      logger.Info,
			Colorful:      true,
		}),
	})
	if err != nil {
		return err
	}
	return nil
}
func GetDefaultDB() (*gorm.DB, error) {
	if DefaultDB == nil {
		panic("db not init")
	}
	return DefaultDB, nil
}

func CloseDB() {

}
