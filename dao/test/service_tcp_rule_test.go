package test

import (
	"gateway/dao"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"reflect"
	"testing"
)

func TestServiceTcpRule_FindOne(t *testing.T) {
	initDB()
	type fields struct {
		ID        uint
		ServiceID uint
		Port      uint16
	}
	type args struct {
		c  *gin.Context
		tx *gorm.DB
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantOut *dao.ServiceTcpRule
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			"查询测试(ID)-记录存在",
			fields{
				ID: 171,
			},
			args{
				c:  &gin.Context{},
				tx: db,
			},
			&dao.ServiceTcpRule{},
			false,
		},
		{
			"查询测试(ID)-记录不存在",
			fields{
				ID: 170,
			},
			args{
				c:  &gin.Context{},
				tx: db,
			},
			&dao.ServiceTcpRule{},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &dao.ServiceTcpRule{
				ServiceID: tt.fields.ServiceID,
				Port:      tt.fields.Port,
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

func TestServiceTcpRule_Save(t *testing.T) {
	type fields struct {
		ID        uint
		ServiceID uint
		Port      uint16
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &dao.ServiceTcpRule{
				Model: gorm.Model{
					ID: tt.fields.ID,
				},
				ServiceID: tt.fields.ServiceID,
				Port:      tt.fields.Port,
			}
			if err := p.Save(tt.args.c, tt.args.tx); (err != nil) != tt.wantErr {
				t.Errorf("Save() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
