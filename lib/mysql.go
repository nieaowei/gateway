package lib

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DefaultDB *gorm.DB

func InitDBPool(path string) error {
	//普通的db方式
	var err error
	DefaultDB, err = gorm.Open(mysql.Open(GetDefaultConfMysql().DataSourceName), &gorm.Config{})
	if err != nil {
		return err
	}
	return nil
}
func GetDefaultDB() (*gorm.DB, error) {
	return DefaultDB, nil
}

func CloseDB() {

}
