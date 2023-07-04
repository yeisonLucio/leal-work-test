package vo

import (
	"reflect"
	"testing"
)

func TestOperator_NewOperator(t *testing.T) {
	type args struct {
		value string
	}
	tests := []struct {
		name    string
		args    args
		want    Operator
		wantErr bool
	}{
		{
			name: "should return an error when type does not exists",
			args: args{
				value: "&",
			},
			wantErr: true,
		},
		{
			name: "should not return an error when type exists",
			args: args{
				value: "%",
			},
			want: Operator{
				value: "%",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewOperator(tt.args.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("Operator.NewOperator() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Operator.NewOperator() = %v, want %v", got, tt.want)
			}
		})
	}
}
