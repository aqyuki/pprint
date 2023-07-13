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
