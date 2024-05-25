package parser

import (
	"github.com/YuukiToriyama/gotiny-wasm-sandbox/types"
	"reflect"
	"testing"
)

func Test_Parse(t *testing.T) {
	type args struct {
		csv string
	}
	tests := []struct {
		name       string
		args       args
		wantResult types.Rows
	}{
		{
			name:       "3列1行",
			args:       args{"id,name,mail_address"},
			wantResult: nil,
		},
		{
			name: "3列2行",
			args: args{"id,name,mail_address\n1,toriyama,toriyama@example.com\n2,hogehoge,hogehoge@example.com"},
			wantResult: types.Rows{
				{"id": "1", "name": "toriyama", "mail_address": "toriyama@example.com"},
				{"id": "2", "name": "hogehoge", "mail_address": "hogehoge@example.com"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := Parse(tt.args.csv); !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("parse() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
