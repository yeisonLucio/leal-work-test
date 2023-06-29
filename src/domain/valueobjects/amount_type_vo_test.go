package valueobjects

import "testing"

func TestAmountType_New(t *testing.T) {
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
			name: "should return an error when type does not exists",
			fields: fields{
				value: "",
			},
			args: args{
				value: "fake",
			},
			wantErr: true,
		},
		{
			name: "should not return error when type is valid",
			fields: fields{
				value: "",
			},
			args: args{
				value: "points",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &AmountType{
				value: tt.fields.value,
			}
			if err := v.New(tt.args.value); (err != nil) != tt.wantErr {
				t.Errorf("AmountType.New() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
