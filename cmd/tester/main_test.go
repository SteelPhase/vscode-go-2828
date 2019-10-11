package main

import "testing"

func Test_doThePrint(t *testing.T) {
	type args struct {
		c config
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "sanity",
			args: args{
				c: config{
					Thing: "12345",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			doThePrint(tt.args.c)
		})
	}
}
