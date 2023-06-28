package valueobjects

import (
	"errors"
	"strconv"
)

type MinAmount struct {
	value float64
}

func (v *MinAmount) NewFromFloat(value float64) {
	v.value = value
}

func (v *MinAmount) NewFromString(value string) error {
	result, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return errors.New("el monto mínimo ingresado no es valido")
	}

	v.value = result

	return nil
}

func (v *MinAmount) GetValue() float64 {
	return v.value
}

func (v *MinAmount) GetStringValue() string {
	return strconv.FormatFloat(v.value, 'f', -1, 64)
}
