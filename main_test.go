package main

import (
	"reflect"
	"testing"
)

func Test_readColumnLengths(t *testing.T) {
	type args struct {
		header string
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "Simple",
			args: args{
				header: "ID   ATTRIBUTE  VALUE",
			},
			want: [][]int{
				{0, 5},
				{5, 16},
				{16, 21},
			},
		},
		{
			name: "Multi-word",
			args: args{
				header: "ID   ATTRIBUTE NAME  CREATED AT",
			},
			want: [][]int{
				{0, 5},
				{5, 21},
				{21, 31},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := readColumnLengths(tt.args.header); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("readColumnLengths() = %v, want %v", got, tt.want)
			}
		})
	}
}
