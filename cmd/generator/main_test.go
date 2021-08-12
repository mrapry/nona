package main

import "testing"

func Test_penjumlahan(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "positif case",
			args: args{
				a: 10,
				b: 5,
			},
			want: 15,
		},
		{
			name: "positif case minus",
			args: args{
				a: -5,
				b: 5,
			},
			want: 0,
		},
		{
			name: "a and b nil",
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := penjumlahan(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("penjumlahan() = %v, want %v", got, tt.want)
			}
		})
	}
}