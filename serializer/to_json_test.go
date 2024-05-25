package serializer

import (
	"github.com/YuukiToriyama/gotiny-wasm-sandbox/types"
	"testing"
)

func Test_ToJson(t *testing.T) {
	type args struct {
		rows types.Rows
	}
	tests := []struct {
		name       string
		args       args
		wantResult string
	}{
		{
			name: "3列2行",
			args: args{rows: types.Rows{
				{"id": "1", "name": "toriyama", "mail_address": "toriyama@example.com"},
				{"id": "2", "name": "hogehoge", "mail_address": "hogehoge@example.com"},
			}},
			wantResult: "[{\"id\": \"1\", \"mail_address\": \"toriyama@example.com\", \"name\": \"toriyama\"}, {\"id\": \"2\", \"mail_address\": \"hogehoge@example.com\", \"name\": \"hogehoge\"}]",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := ToJson(tt.args.rows); gotResult != tt.wantResult {
				t.Errorf("ToJson() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
