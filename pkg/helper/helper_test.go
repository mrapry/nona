package helper

import (
	"reflect"
	"testing"
)

func TestCheckStringNotNull(t *testing.T) {
	type args struct {
		data string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Positif case",
			args: args{data: "oke"},
			want: true,
		},
		{
			name: "negatif case",
			args: args{},
			want: false,
		},

	}
		for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CheckStringNotNull(tt.args.data); got != tt.want {
				t.Errorf("CheckStringNotNull() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringInSlice(t *testing.T) {
	type args struct {
		str  string
		list []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "positif",
			args: args{
				str:  "dua",
				list: []string{"satu","dua"},
			},
			want: true,
		},
		{
			name: "negatif",
			args: args{
				str:  "dua",
				list: []string{"satu","tiga"},
			},
			want: false,
		},
	}
		for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StringInSlice(tt.args.str, tt.args.list); got != tt.want {
				t.Errorf("StringInSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToBytes(t *testing.T) {
	type args struct {
		i interface{}
	}
	tests := []struct {
		name  string
		args  args
		wantB []byte
	}{
		{
			name:  "positif",
			args:  args{"halo"},
			wantB: ToBytes("halo"),
		},
		{
			name:  "positif",
			args:  args{ToBytes("halo")},
			wantB: ToBytes("halo"),
		},
		{
			name:  "positif",
			args:  args{123},
			wantB: ToBytes(123),
		},
	}
		for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotB := ToBytes(tt.args.i); !reflect.DeepEqual(gotB, tt.wantB) {
				t.Errorf("ToBytes() = %v, want %v", gotB, tt.wantB)
			}
		})
	}
}