package valueobjects

import (
	"errors"
)

type AmountType struct {
	value string
}

var types = []string{"coins", "points"}

func (v *AmountType) New(value string) error {
	found := false
	for _, v := range types {
		if v == value {
			found = true
			break
		}
	}

	if !found {
		return errors.New("el tipo de monto ingresado no es valido")
	}

	v.value = value

	return nil
}

func (v *AmountType) GetValue() string {
	return v.value
}
