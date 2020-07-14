package dao

import (
	"github.com/e421083458/golang_common/lib"
	"github.com/e421083458/gorm"
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"reflect"
	"testing"
)

var db *gorm.DB

func initDB() {
	lib.InitDBPool("../conf/dev/mysql_map.toml")
	db, _ = lib.GetGormPool("default")
	db.SetLogger(gorm.Logger{log.New(os.Stdout, "\n", 0)})
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
				tx: db.Debug(),
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

func TestAdmin_Updates(t *testing.T) {
	initDB()
	type fields struct {
		Model    gorm.Model
		Salt     string
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
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			"空值测试-1",
			fields{
				Model: gorm.Model{
					ID: 1,
				},
				Username: "admin",
				Salt:     "123",
			},
			args{
				c:  &gin.Context{},
				tx: db,
			},
			false,
		},
		{
			"空值测试-2",
			fields{
				Model: gorm.Model{
					ID: 1,
				},
				Username: "admin",
				Salt:     "",
			},
			args{
				c:  &gin.Context{},
				tx: db,
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Admin{
				Model:    tt.fields.Model,
				Salt:     tt.fields.Salt,
				Username: tt.fields.Username,
				Password: tt.fields.Password,
			}
			if err := p.Updates(tt.args.c, tt.args.tx); (err != nil) != tt.wantErr {
				t.Errorf("Updates() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestAdmin_Save(t *testing.T) {
	initDB()
	type fields struct {
		Model    gorm.Model
		Salt     string
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
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			"空值更新测试-1",
			fields{
				Model: gorm.Model{
					ID: 1,
				},
				Password: "6d8b2aadeecc1a9504b396ad74697f5675aca7d6751c42747ac42403cb3b9ef7",
				Salt:     "123",
			},
			args{
				c:  &gin.Context{},
				tx: db,
			},
			false,
		},
		{
			"空值更新测试-1",
			fields{
				Model: gorm.Model{
					ID: 1,
				},
				Password: "6d8b2aadeecc1a9504b396ad74697f5675aca7d6751c42747ac42403cb3b9ef7",
				Salt:     "123",
			},
			args{
				c:  &gin.Context{},
				tx: db,
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Admin{
				Model:    tt.fields.Model,
				Salt:     tt.fields.Salt,
				Username: tt.fields.Username,
				Password: tt.fields.Password,
			}
			if err := p.Save(tt.args.c, tt.args.tx); (err != nil) != tt.wantErr {
				t.Errorf("Save() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
