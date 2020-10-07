package lib

import (
	"gorm.io/gorm"
	"testing"
)

func TestGetDefaultDB(t *testing.T) {
	tests := []struct {
		name   string
		wantDb *gorm.DB
	}{
		// TODO: AddHost test cases.
		{
			name:   "",
			wantDb: &gorm.DB{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotDb := GetDefaultDB(); gotDb == nil {
				t.Errorf("GetDefaultDB() = %v, want %v", gotDb, tt.wantDb)
			}
		})
	}
}

func TestInitDBPool(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: AddHost test cases.
		{
			name: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			InitDBPool()
			if DefaultDB == nil {
				t.Errorf("DefaultDB %v", DefaultDB)
			}
		})
	}
}

func BenchmarkGetDefaultDB(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GetDefaultDB()
	}
}
