package vo

import (
	"errors"

	"lucio.com/order-service/src/domain/helpers"
)

type AmountType struct {
	value string
}

var amountTypes = []string{"coins", "points"}

func NewAmountType(value string) (AmountType, error) {
	amountType := AmountType{}
	if !helpers.StringContains(amountTypes, value) {
		return amountType, errors.New("el tipo de monto ingresado no es valido")
	}

	amountType.value = value

	return amountType, nil
}

func (v *AmountType) Value() string {
	return v.value
}
