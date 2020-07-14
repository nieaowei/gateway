package dao

import (
	"github.com/e421083458/gorm"
	"github.com/gin-gonic/gin"
	"reflect"
	"testing"
)

func TestServiceInfo_PageList(t *testing.T) {
	initDB()
	type fields struct {
		Model       gorm.Model
		LoadType    uint
		ServiceName string
		ServiceDesc string
	}
	type args struct {
		c      *gin.Context
		tx     *gorm.DB
		params *PageSize
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		wantList  []ServiceInfo
		wantCount uint
		wantErr   bool
	}{
		// TODO: Add test cases.
		{
			"查询测试-1",
			fields{},
			args{
				c:  &gin.Context{},
				tx: db,
				params: &PageSize{
					Size: 2,
					No:   1,
					Info: "",
				},
			},
			nil,
			2,
			false,
		},
		{
			"查询测试-2",
			fields{},
			args{
				c:  &gin.Context{},
				tx: db,
				params: &PageSize{
					Size: 2,
					No:   2,
					Info: "",
				},
			},
			nil,
			2,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &ServiceInfo{
				Model:       tt.fields.Model,
				LoadType:    tt.fields.LoadType,
				ServiceName: tt.fields.ServiceName,
				ServiceDesc: tt.fields.ServiceDesc,
			}
			gotList, gotCount, err := p.PageList(tt.args.c, tt.args.tx, tt.args.params)
			if !reflect.DeepEqual(gotList, tt.wantList) {
				t.Logf("PageList() gotList = %v, want %v", gotList, tt.wantList)
			}
			if gotCount != tt.wantCount {
				t.Logf("PageList() gotCount = %v, want %v", gotCount, tt.wantCount)
			}
			if (err != nil) != tt.wantErr {
				t.Errorf("PageList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

		})
	}
}
