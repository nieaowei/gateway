package public

import "testing"

func TestGenSha256BySecret(t *testing.T) {
	type args struct {
		str    string
		secret string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: AddHost test cases.
		{
			"生成测试-1",
			args{
				secret: "123",
				str:    "nieaowei",
			},
			"2133e596930eba812c573636a5e6aa5686cb7ced608a7025c44980dc03fb2f48",
		},
		{
			"生成测试-2",
			args{
				secret: "1234",
				str:    "nieaowei",
			},
			"13d125c12461ae5de6c296f836005a2ace91d79d37810417788065b1e756043a",
		},
		{
			"生成测试-3",
			args{
				secret: "as",
				str:    "nieaowei",
			},
			"c350c3edbe40ab228c4e54e33e99678e86bb11d711d73822319491c8e4fa522d",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GenSha256BySecret(tt.args.str, tt.args.secret)
			t.Logf("GenSha256BySecret() = %v", got)
			if got != tt.want {
				t.Errorf("GenSha256BySecret() = %v, want %v", got, tt.want)
			}
		})
	}
}
