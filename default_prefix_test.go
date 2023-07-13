package pprint_test

import (
	"reflect"
	"testing"

	"github.com/aqyuki/pprint"
)

func TestNew(t *testing.T) {
	type args struct {
		word string
	}
	tests := []struct {
		name string
		args args
		want *pprint.DefaultPrefix
	}{
		{
			name: "case 1",
			args: args{
				word: "Error",
			},
			want: &pprint.DefaultPrefix{
				Message: "Error",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pprint.New(tt.args.word); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGet(t *testing.T) {
	type args struct {
		prefix string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "case 1",
			args: args{
				prefix: "Error",
			},
			want: "[Error]",
		},
	}

	for _, tt := range tests {
		prefix := pprint.DefaultPrefix{
			Message: tt.args.prefix,
		}
		t.Run(tt.name, func(t *testing.T) {
			if got := prefix.Get(); tt.want != got {
				t.Errorf("DefaultPrefix.Get() => got : %v , want : %v", got, tt.want)
			}
		})
	}
}
