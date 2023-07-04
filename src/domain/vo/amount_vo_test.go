package vo

import (
	"reflect"
	"testing"
)

func TestAmount_NewAmountFromString(t *testing.T) {
	type args struct {
		value string
	}
	tests := []struct {
		name    string
		args    args
		want    Amount
		wantErr bool
	}{
		{
			name: "should return an error when value does not float",
			args: args{
				value: "fake",
			},
			wantErr: true,
		},
		{
			name: "should not return an error when value is a number valid",
			args: args{
				value: "2",
			},
			want: Amount{
				value: 2,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewAmountFromString(tt.args.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("Amount.NewAmountFromString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Amount.NewAmountFromString() = %v, want %v", got, tt.want)
			}
		})
	}
}
