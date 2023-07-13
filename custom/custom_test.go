package custom

import "testing"

func TestCustomPrefix_Get(t *testing.T) {
	tests := []struct {
		name string
		p    CustomPrefix
		want string
	}{
		{
			name: "case 1",
			p: CustomPrefix{
				Builder: func() string { return "[sample]" },
			},
			want: "[sample]",
		},
		{
			name: "case 2",
			p: CustomPrefix{
				Builder: func() string { return "sample:" },
			},
			want: "sample:",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.Get(); got != tt.want {
				t.Errorf("CustomPrefix.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}
