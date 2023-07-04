package vo

import (
	"errors"
	"strconv"
)

type Amount struct {
	value float64
}

func NewAmountFromFloat(value float64) Amount {
	amount := Amount{
		value: value,
	}

	return amount
}

func NewAmountFromString(value string) (Amount, error) {
	amount := Amount{}
	result, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return amount, errors.New("el monto ingresado no es valido")
	}

	amount.value = result

	return amount, nil
}

func (v *Amount) Value() float64 {
	return v.value
}

func (v *Amount) String() string {
	return strconv.FormatFloat(v.value, 'f', -1, 64)
}
