package valueobjects

import "testing"

func TestAmount_NewFromString(t *testing.T) {
	type fields struct {
		value float64
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
			name:   "should return an error when value does not float",
			fields: fields{},
			args: args{
				value: "fake",
			},
			wantErr: true,
		},
		{
			name:   "should not return an error when value is a number valid",
			fields: fields{},
			args: args{
				value: "2",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &Amount{
				value: tt.fields.value,
			}
			if err := v.NewFromString(tt.args.value); (err != nil) != tt.wantErr {
				t.Errorf("Amount.NewFromString() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
