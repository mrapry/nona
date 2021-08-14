package helper

import (
	"reflect"
	"testing"
	"time"
)

func TestConvertStringToTimeStamp(t *testing.T) {
	type args struct {
		dataTime string
	}
	tests := []struct {
		name string
		args args
		want time.Time
	}{
		{
			name: "Positif test",
			args: args{dataTime: "2021-01-01"},
			want: ConvertStringToTimeStamp("2021-01-01"),
		},
		{
			name: "Negatif test",
			args: args{dataTime: "2021023"},
			want: time.Time{},
		},
	}
		for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ConvertStringToTimeStamp(tt.args.dataTime); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ConvertStringToTimeStamp() = %v, want %v", got, tt.want)
			}
		})
	}
}
