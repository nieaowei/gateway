package dao

import (
	"errors"
	"gorm.io/gorm"
	"reflect"
)

const (
	LoadTypeHttp = 0
	LoadTypeTcp  = 1
	LoadTypeGrpc = 2

	HttpRuleTypePrefixURL = 0
	HttpRuleTypeDomain    = 1
)

type PageSize struct {
	Size int
	No   int
	Info string
}

func GetFields(p interface{}) []string {
	t := reflect.TypeOf(p).Elem()
	all := []string{}
	for i := 0; i < t.NumField(); i++ {
		if t.Field(i).Type == reflect.TypeOf(gorm.Model{}) {
			for j := 0; j < t.Field(i).Type.NumField(); j++ {
				all = append(all, t.Field(i).Type.Field(j).Name)
			}
			continue
		}
		all = append(all, t.Field(i).Name)
	}
	return all
}

func StructToMap(obj interface{}) map[string]interface{} {
	t := reflect.TypeOf(obj).Elem()
	v := reflect.ValueOf(obj).Elem()

	var data = make(map[string]interface{})

	for i := 0; i < t.NumField(); i++ {
		data[t.Field(i).Name] = v.Field(i).Interface()
	}
	return data
}

func ErrorHandleForDB(res *gorm.DB) (err error) {
	if res.Error != nil {
		err = res.Error
		return
	}
	if res.RowsAffected == 0 {
		err = errors.New("not updated")
		return
	}
	return
}
