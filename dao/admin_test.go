package dao

import (
	"gateway/lib"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"testing"
)

func TestAdmin_FindOneScan(t *testing.T) {
	type fields struct {
		Model    gorm.Model
		Username string
		Salt     string
		Password string
	}
	type args struct {
		c   *gin.Context
		tx  *gorm.DB
		out interface{}
	}
	type out struct {
		ID       uint `json:"id"`
		Username string
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
				Username: "admin",
			},
			args: args{
				c:   nil,
				tx:  lib.GetDefaultDB(),
				out: &out{},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Admin{
				Model:    tt.fields.Model,
				Username: tt.fields.Username,
				Salt:     tt.fields.Salt,
				Password: tt.fields.Password,
			}
			err := p.FindOneScan(tt.args.c, tt.args.tx, tt.args.out)
			t.Log(tt.args.out)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindOneScan() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestAdmin_UpdateByID(t *testing.T) {
	type fields struct {
		Model    gorm.Model
		Username string
		Salt     string
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Admin{
				Model:    tt.fields.Model,
				Username: tt.fields.Username,
				Salt:     tt.fields.Salt,
				Password: tt.fields.Password,
			}
			if err := p.UpdateByID(tt.args.c, tt.args.tx); (err != nil) != tt.wantErr {
				t.Errorf("UpdateByID() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func BenchmarkAdmin_FindOne(b *testing.B) {
	admin := &Admin{
		Username: "admin",
	}
	db := lib.GetDefaultDB()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = admin.FindOne(nil, db)
	}
}

func TestAdmin_UpdateAllByID(t *testing.T) {
	type fields struct {
		Model    gorm.Model
		Username string
		Salt     string
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
			name: "",
			fields: fields{
				Username: "nieaowei",
				Salt:     "123",
				Password: "nieaowei123",
			},
			args: args{
				c:  nil,
				tx: lib.GetDefaultDB(),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Admin{
				Model:    tt.fields.Model,
				Username: tt.fields.Username,
				Salt:     tt.fields.Salt,
				Password: tt.fields.Password,
			}
			if err := p.UpdateAllByID(tt.args.c, tt.args.tx); (err != nil) != tt.wantErr {
				t.Errorf("UpdateAllByID() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
