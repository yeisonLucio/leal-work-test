package valueobjects

import (
	"errors"

	"lucio.com/order-service/src/domain/helpers"
)

type Operation struct {
	value string
}

var operatorTypes = []string{"%", "*"}

func (v *Operation) New(value string) error {
	if !helpers.StringContains(operatorTypes, value) {
		return errors.New("el operador ingresado no es valido")
	}

	v.value = value

	return nil
}

func (v *Operation) GetValue() string {
	return v.value
}
