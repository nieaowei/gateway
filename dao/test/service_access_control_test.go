package test

import (
	"gateway/dao"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"testing"
)

func TestServiceAccessControl_UpdateAll(t *testing.T) {
	initDB()
	type fields struct {
		Model             gorm.Model
		ServiceID         uint
		OpenAuth          int8
		BlackList         string
		WhiteList         string
		WhiteHostName     string
		ClientipFlowLimit int
		ServiceFlowLimit  int
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
				ServiceID:        97,
				ServiceFlowLimit: 100,
			},
			args: args{
				db: db,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &dao.ServiceAccessControl{
				Model:             tt.fields.Model,
				ServiceID:         tt.fields.ServiceID,
				OpenAuth:          tt.fields.OpenAuth,
				BlackList:         tt.fields.BlackList,
				WhiteList:         tt.fields.WhiteList,
				WhiteHostName:     tt.fields.WhiteHostName,
				ClientipFlowLimit: tt.fields.ClientipFlowLimit,
				ServiceFlowLimit:  tt.fields.ServiceFlowLimit,
			}
			if err := p.UpdateAllByServiceID(tt.args.c, tt.args.db); (err != nil) != tt.wantErr {
				t.Errorf("UpdateAllByServiceID() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
