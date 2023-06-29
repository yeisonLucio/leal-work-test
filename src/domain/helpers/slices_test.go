package helpers

import "testing"

func TestStringContains(t *testing.T) {
	type args struct {
		options []string
		value   string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "should return true when item exists",
			args: args{
				options: []string{"one", "two", "three"},
				value:   "two",
			},
			want: true,
		},
		{
			name: "should return false when item does not exists",
			args: args{
				options: []string{"one", "two", "three"},
				value:   "four",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StringContains(tt.args.options, tt.args.value); got != tt.want {
				t.Errorf("StringContains() = %v, want %v", got, tt.want)
			}
		})
	}
}
