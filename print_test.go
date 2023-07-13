package pprint_test

import (
	"testing"

	"github.com/aqyuki/pprint"
	"github.com/aqyuki/pprint/custom"
)

func TestErrorPrint(t *testing.T) {

	type args struct {
		message string
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "case 1",
			args: args{
				message: "sample",
			},
			want: "[Error] sample",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := pprint.PickStdout(t, func() { pprint.ErrorPrint(tt.args.message) })
			if got != tt.want {
				t.Errorf("output : %v , want : %v\n", got, tt.want)
			}
		})
	}
}
func TestInfoPrint(t *testing.T) {

	type args struct {
		message string
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "case 1",
			args: args{
				message: "sample",
			},
			want: "[Info] sample",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := pprint.PickStdout(t, func() { pprint.InfoPrint(tt.args.message) })
			if got != tt.want {
				t.Errorf("output : %v , want : %v\n", got, tt.want)
			}
		})
	}
}

func TestCustomPrint(t *testing.T) {
	type args struct {
		message string
		prefix  custom.CustomPrefix
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "case 1",
			args: args{
				message: "sample",
				prefix: custom.CustomPrefix{
					Builder: func() string {
						return "info:"
					},
				},
			},
			want: "info: sample",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := pprint.PickStdout(t, func() { pprint.CustomPrint(tt.args.message, tt.args.prefix) })
			if got != tt.want {
				t.Errorf("output : %v , want : %v\n", got, tt.want)
			}
		})
	}
}
