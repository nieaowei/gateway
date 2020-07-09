package public

import (
	"fmt"
	"testing"
)

func TestGenSha256ByScret(t *testing.T) {
	type args struct {
		scret string
		str   string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			"生成测试-1",
			args{
				scret: "123",
				str:   "nieaowei",
			},
			"732e085daa85ca813a29294994e96af64d9f644ff41a6d6cbcd62c64025e2f78",
		},
		{
			"生成测试-2",
			args{
				scret: "1234",
				str:   "nieaowei",
			},
			"a45557db71b088ea9f177f2884701a04e9f3faedb3f6a27bd96759e4baf84752",
		},
		{
			"生成测试-3",
			args{
				scret: "as",
				str:   "nieaowei",
			},
			"9dc53c8fd9ecab422c0eb8d578853cdbcddc5b4222a9d3f6986581a022f6bd66",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GenSha256ByScret(tt.args.scret, tt.args.str);
			fmt.Println("result: "+got)
			if  got != tt.want {
				t.Errorf("GenSha256ByScret() = %v, want %v", got, tt.want)
			}
		})
	}
}
