package test

import (
	"gateway/dao"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"reflect"
	"testing"
)

func TestServiceInfo_PageList(t *testing.T) {
	initDB()
	type fields struct {
		Model       gorm.Model
		LoadType    int8
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
		wantCount int64
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

func TestServiceInfo_AddAfterCheck(t *testing.T) {
	initDB()
	type fields struct {
		Model       gorm.Model
		LoadType    int8
		ServiceName string
		ServiceDesc string
	}
	type args struct {
		c     *gin.Context
		db    *gorm.DB
		check bool
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "1",
			fields: fields{
				Model:       gorm.Model{},
				LoadType:    0,
				ServiceName: "11433323",
				ServiceDesc: "123",
			},
			args: args{
				c:  &gin.Context{},
				db: db,
			},
			wantErr: false,
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
			if err := p.InsertAfterCheck(tt.args.c, tt.args.db, tt.args.check); (err != nil) != tt.wantErr {
				t.Errorf("InsertAfterCheck() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestServiceInfo_Delete(t *testing.T) {
	initDB()
	type fields struct {
		Model       gorm.Model
		LoadType    int8
		ServiceName string
		ServiceDesc string
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
			name: "",
			fields: fields{
				Model: gorm.Model{
					ID: 95,
				},
			},
			args: args{
				tx: db,
			},
			wantErr: false,
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
			if err := p.DeleteByID(tt.args.c, tt.args.tx); (err != nil) != tt.wantErr {
				t.Errorf("DeleteByID() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestServiceInfo_UpdateAll(t *testing.T) {
	initDB()
	type fields struct {
		Model       gorm.Model
		LoadType    int8
		ServiceName string
		ServiceDesc string
	}
	type args struct {
		c  *gin.Context
		db *gorm.DB
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "",
			fields: fields{
				Model: gorm.Model{
					ID: 96,
				},
				ServiceName: "33344",
			},
			args: args{
				db: db,
			},
			wantErr: false,
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
			if err := p.UpdateAll(tt.args.c, tt.args.db); (err != nil) != tt.wantErr {
				t.Errorf("UpdateAllByServiceID() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
