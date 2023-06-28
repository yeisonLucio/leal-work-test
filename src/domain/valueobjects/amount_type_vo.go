package valueobjects

import (
	"errors"

	"lucio.com/order-service/src/domain/helpers"
)

type AmountType struct {
	value string
}

var amountTypes = []string{"coins", "points"}

func (v *AmountType) New(value string) error {
	if !helpers.StringContains(amountTypes, value) {
		return errors.New("el tipo de monto ingresado no es valido")
	}

	v.value = value

	return nil
}

func (v *AmountType) GetValue() string {
	return v.value
}
