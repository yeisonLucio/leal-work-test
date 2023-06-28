package valueobjects

import (
	"errors"
	"strconv"
)

type Amount struct {
	value float64
}

func (v *Amount) NewFromFloat(value float64) {
	v.value = value
}

func (v *Amount) NewFromString(value string) error {
	result, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return errors.New("el monto ingresado no es valido")
	}

	v.value = result

	return nil
}

func (v *Amount) GetValue() float64 {
	return v.value
}

func (v *Amount) GetStringValue() string {
	return strconv.FormatFloat(v.value, 'f', -1, 64)
}
