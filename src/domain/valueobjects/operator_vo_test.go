package valueobjects

import "testing"

func TestOperator_New(t *testing.T) {
	type fields struct {
		value string
	}
	type args struct {
		value string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:   "should return an error when type does not exists",
			fields: fields{},
			args: args{
				value: "&",
			},
			wantErr: true,
		},
		{
			name:   "should not return an error when type exists",
			fields: fields{},
			args: args{
				value: "%",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &Operator{
				value: tt.fields.value,
			}
			if err := v.New(tt.args.value); (err != nil) != tt.wantErr {
				t.Errorf("Operator.New() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
