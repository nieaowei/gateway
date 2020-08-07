package lib

import (
	"testing"
)

func TestGetDefaultConfBase(t *testing.T) {
	tests := []struct {
		name string
		want *BaseConf
	}{
		// TODO: Add test cases.
		{
			name: "",
			want: &BaseConf{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetDefaultConfBase(); got == nil {
				t.Errorf("GetDefaultConfBase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetDefaultConfMysql(t *testing.T) {
	tests := []struct {
		name string
		want *MySQLConf
	}{
		// TODO: Add test cases.
		{
			name: "",
			want: &MySQLConf{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetDefaultConfMysql(); got == nil {
				t.Errorf("GetDefaultConfMysql() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetDefaultConfRedis(t *testing.T) {
	tests := []struct {
		name string
		want *RedisConf
	}{
		// TODO: Add test cases.
		{
			name: "",
			want: &RedisConf{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetDefaultConfRedis(); got == nil {
				t.Errorf("GetDefaultConfRedis() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkGetDefaultConfBase(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GetDefaultConfBase()
	}
}

func BenchmarkGetDefaultConfRedis(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GetDefaultConfRedis()
	}
}

func BenchmarkGetDefaultConfMysql(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GetDefaultConfMysql()
	}
}
