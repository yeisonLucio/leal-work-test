package entities

import (
	"errors"
	"time"

	"lucio.com/order-service/src/domain/vo"
)

var (
	errStartDateInvalid = errors.New("el formato de start_date es incorrecto")
	errEndDateInvalid   = errors.New("el formato de end_date es incorrecto")
)

type BranchCampaign struct {
	ID            uint
	CampaignID    uint
	BranchID      uint
	StartDate     time.Time
	EndDate       time.Time
	Operator      vo.Operator
	OperatorValue uint
	MinAmount     vo.Amount
}

func (b *BranchCampaign) SetStartDate(date string) error {
	time, err := time.Parse(time.DateTime, date)
	if err != nil {
		return errStartDateInvalid
	}

	b.StartDate = time
	return nil
}

func (b *BranchCampaign) SetEndDate(date string) error {
	time, err := time.Parse(time.DateTime, date)
	if err != nil {
		return errEndDateInvalid
	}

	b.EndDate = time
	return nil
}

func (b *BranchCampaign) SetOperator(value string) error {
	operator, err := vo.NewOperator(value)
	if err != nil {
		return err
	}

	b.Operator = operator
	return nil
}
