package test

import (
	"gateway/dao"
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
		params *dao.PageSize
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		wantList  []dao.ServiceInfo
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
				params: &dao.PageSize{
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
				params: &dao.PageSize{
					Size: 2,
					No:   2,
					Info: "",
				},
			},
			nil,
			2,
			false,
		},
		{
			"查询测试-3",
			fields{},
			args{
				c:  &gin.Context{},
				tx: db,
				params: &dao.PageSize{
					Size: 22,
					No:   1,
					Info: "tcp",
				},
			},
			nil,
			2,
			false,
		},
		{
			"查询测试-3",
			fields{},
			args{
				c:  &gin.Context{},
				tx: db,
				params: &dao.PageSize{
					Size: 20,
					No:   1,
					Info: "udp",
				},
			},
			nil,
			2,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &dao.ServiceInfo{
				Model:       tt.fields.Model,
				LoadType:    tt.fields.LoadType,
				ServiceName: tt.fields.ServiceName,
				ServiceDesc: tt.fields.ServiceDesc,
			}
			gotList, gotCount, err := p.PageList(tt.args.c, tt.args.tx, tt.args.params)
			if !reflect.DeepEqual(gotList, tt.wantList) {
				t.Logf("PageListScan() gotList = %v, want %v", gotList, tt.wantList)
			}
			if gotCount != tt.wantCount {
				t.Logf("PageListScan() gotCount = %v, want %v", gotCount, tt.wantCount)
			}
			if (err != nil) != tt.wantErr {
				t.Errorf("PageListScan() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

		})
	}
}
