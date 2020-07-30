package lib

import "testing"

func TestInitBaseConf(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "",
			args: args{
				path: "../conf/dev",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := InitBaseConf(tt.args.path); (err != nil) != tt.wantErr {
				t.Errorf("InitBaseConf() error = %v, wantErr %v", err, tt.wantErr)
			}
			t.Log(GetDefaultConfBase())
		})
	}
}

func TestInitRedisConf(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "",
			args: args{
				path: "../conf/dev",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := InitRedisConf(tt.args.path); (err != nil) != tt.wantErr {
				t.Errorf("InitRedisConf() error = %v, wantErr %v", err, tt.wantErr)
			}
			t.Log(GetDefaultConfRedis())
		})
	}
}

func TestInitMysqlConf(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "",
			args: args{
				path: "../conf/dev",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := InitMysqlConf(tt.args.path); (err != nil) != tt.wantErr {
				t.Errorf("InitMysqlConf() error = %v, wantErr %v", err, tt.wantErr)
			}
			t.Log(GetDefaultConfMysql())
		})
	}
}
