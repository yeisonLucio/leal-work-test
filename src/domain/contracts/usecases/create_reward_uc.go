package usecases

import "lucio.com/order-service/src/domain/dto"

type CreateRewardUC interface {
	Execute(createBranchDTO dto.CreateRewardDTO) (*dto.RewardCreatedDTO, error)
}
