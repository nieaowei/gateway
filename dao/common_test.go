package dao

import (
	"testing"
)

type Test struct {
	Name string
	Age  int
	Ip   string
}

func TestGetFields(t *testing.T) {
	type args struct {
		p interface{}
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// TODO: AddHost test cases.
		{
			name: "",
			args: args{
				p: &Test{},
			},
			want: []string{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetFields(tt.args.p); got == nil {
				t.Errorf("GetFields() = %v,", got)
			}
		})
	}
}

func TestStructToMap(t *testing.T) {
	type args struct {
		obj interface{}
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		// TODO: AddHost test cases.
		{
			name: "",
			args: args{
				obj: &Test{},
			},
			want: map[string]interface{}{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StructToMap(tt.args.obj); got == nil {
				t.Errorf("StructToMap() = %v,", got)
			}
		})
	}
}

func BenchmarkGetFields(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = GetFields(&Test{})
	}
}

func BenchmarkStructToMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = StructToMap(&Test{})
	}
}
