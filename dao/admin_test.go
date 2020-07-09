package dao

import (
	"github.com/e421083458/golang_common/lib"
	"github.com/e421083458/gorm"
	"github.com/gin-gonic/gin"
	"reflect"
	"testing"
)

var db *gorm.DB

func initDB() {
	lib.InitDBPool("../conf/dev/mysql_map.toml")
	db, _ = lib.GetGormPool("default")
}

func TestAdmin_FindOne(t *testing.T) {
	initDB()
	type fields struct {
		Model    gorm.Model
		Username string
		Password string
	}
	type args struct {
		c  *gin.Context
		tx *gorm.DB
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantOut *Admin
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			"查询测试(ID)",
			fields{
				Model: gorm.Model{
					ID: 1,
				},
			},
			args{
				c:  &gin.Context{},
				tx: db,
			},
			&Admin{},
			false,
		},
		{
			"查询测试(Username)",
			fields{
				Username: "admin",
			},
			args{
				c:  &gin.Context{},
				tx: db,
			},
			&Admin{},
			false,
		},
		{
			"查询测试(空)",
			fields{
				Username: "1",
			},
			args{
				c:  &gin.Context{},
				tx: db,
			},
			&Admin{},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Admin{
				Model:    tt.fields.Model,
				Username: tt.fields.Username,
				Password: tt.fields.Password,
			}
			gotOut, err := p.FindOne(tt.args.c, tt.args.tx)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindOne() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotOut, tt.wantOut) {
				t.Logf("FindOne() gotOut = %v, want %v", gotOut, tt.wantOut)
			}
		})
	}
}
