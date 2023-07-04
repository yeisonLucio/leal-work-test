package entities

import (
	"time"

	"lucio.com/order-service/src/domain/vo"
)

type BranchCampaign struct {
	ID             uint
	CampaignID     uint
	BranchID       uint
	StartDate      time.Time
	EndDate        time.Time
	Operator       vo.Operator
	OperationValue uint
	MinAmount      vo.Amount
}
