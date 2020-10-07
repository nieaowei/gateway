package loadbalance

import (
	"net/url"
	"strconv"
	"testing"
)

func TestWeightRobinLoadBalance_Get(t *testing.T) {
	type fields struct {
		hostList      []*WeightNode
		currentWeight int
	}
	type args struct {
		key string
	}
	data := make([]*url.URL, 5)
	for i := 0; i < 5; i++ {
		data[i], _ = url.Parse("http://127.0.0.1:" + strconv.Itoa(i))
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		// TODO: AddHost test cases.
		{
			name: "",
			fields: fields{
				hostList: []*WeightNode{
					{
						Weight:        7,
						CurrentWeight: 0,
						Addr:          data[0],
					},
					{
						Weight:        7,
						CurrentWeight: 0,
						Addr:          data[1],
					},
					{
						Weight:        7,
						CurrentWeight: 0,
						Addr:          data[2],
					},
					{
						Weight:        7,
						CurrentWeight: 0,
						Addr:          data[3],
					},
					{
						Weight:        12,
						CurrentWeight: 0,
						Addr:          data[4],
					},
				},
				currentWeight: 0,
			},
			args: args{
				key: "sss",
			},
			want:    "",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &WeightRobinLoadBalance{
				hostList:      tt.fields.hostList,
				currentWeight: tt.fields.currentWeight,
			}
			for i := 0; i < 10; i++ {
				got, err := w.GetHost(tt.args.key)
				if (err != nil) != tt.wantErr {
					t.Errorf("GetHost() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				//assert.NotEqual(t,got,"")
				if got == nil {
					t.Errorf("GetHost() got = %v, want %v", got, tt.want)
				}
			}

		})
	}
}
