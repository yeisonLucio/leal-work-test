package vo

import (
	"reflect"
	"testing"
)

func TestAmountType_NewAmountType(t *testing.T) {
	type args struct {
		value string
	}
	tests := []struct {
		name    string
		args    args
		want    AmountType
		wantErr bool
	}{
		{
			name: "should return an error when type does not exists",
			args: args{
				value: "fake",
			},
			wantErr: true,
		},
		{
			name: "should not return error when type is valid",
			args: args{
				value: "points",
			},
			want: AmountType{
				value: "points",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewAmountType(tt.args.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("AmountType.NewAmountType() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AmountType.NewAmountType() = %v, want %v", got, tt.want)
			}
		})
	}
}
