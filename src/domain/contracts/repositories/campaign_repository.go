package repositories

import "lucio.com/order-service/src/domain/entities"

type CampaignRepository interface {
	Create(campaign entities.Campaign) (*entities.Campaign, error)
	FindByID(ID uint) *entities.Campaign
}
