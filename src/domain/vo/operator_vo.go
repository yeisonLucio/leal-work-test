package vo

import (
	"errors"

	"lucio.com/order-service/src/domain/helpers"
)

type Operator struct {
	value string
}

var operatorTypes = []string{"%", "*"}

func NewOperator(value string) (Operator, error) {
	operator := Operator{}
	if !helpers.StringContains(operatorTypes, value) {
		return operator, errors.New("el operador ingresado no es valido")
	}

	operator.value = value

	return operator, nil
}

func (v *Operator) Value() string {
	return v.value
}
