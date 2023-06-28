package entities

import (
	"time"

	"lucio.com/order-service/src/domain/valueobjects"
)

type BranchCampaign struct {
	ID             uint
	CampaignID     uint
	BranchID       uint
	StartDate      time.Time
	EndDate        time.Time
	Operator       valueobjects.Operation
	OperationValue uint
	MinAmount      valueobjects.Amount
}
