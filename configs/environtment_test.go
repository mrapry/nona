package configs

import (
	"os"
	"reflect"
	"testing"
)

func TestGetEnv(t *testing.T) {
	tests := []struct {
		name string
		want Environment
	}{
		{
			name: "positifCase",
			want: GetEnv(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetEnv(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetEnv() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_loadAdditionalEnv(t *testing.T) {
	tests := []struct {
		name string
		envValue struct{
			key string
			value string
		}
		wantError bool
	}{
		{
			name: "positif",
			envValue: struct {
				key   string
				value string
			}{key: "sample", value: "123"},
			wantError: false,
		},
	}
	for _, tt := range tests {
		os.Setenv(tt.envValue.key, tt.envValue.value)
		t.Run(tt.name, func(t *testing.T) {
			if !tt.wantError {
				assertPanic(t, loadAdditionalEnv)
			}
		})
	}
}

func assertPanic(t *testing.T, f func()) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	f()
}
