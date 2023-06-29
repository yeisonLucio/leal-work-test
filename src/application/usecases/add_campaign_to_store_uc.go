package usecases

import (
	"errors"
	"sync"

	"lucio.com/order-service/src/domain/contracts/repositories"
	"lucio.com/order-service/src/domain/contracts/usecases"
	"lucio.com/order-service/src/domain/dto"
)

type AddCampaignToStoreUC struct {
	BranchRepository       repositories.BranchRepository
	CreateBranchCampaignUC usecases.CreateBranchCampaignUC
	CampaignRepository     repositories.CampaignRepository
	StoreRepository        repositories.StoreRepository
}

type resultProcess struct {
	err      error
	result   *dto.BranchCampaignCreatedDTO
	BranchID uint
}

func (a *AddCampaignToStoreUC) Execute(
	createStoreCampaignDTO dto.CreateStoreCampaignDTO,
) (*dto.StoreCampaignCreatedDTO, error) {
	if campaign := a.CampaignRepository.FindByID(createStoreCampaignDTO.CampaignID); campaign == nil {
		return nil, errors.New("la campaña no existe")
	}

	if store := a.StoreRepository.FindByID(createStoreCampaignDTO.StoreID); store == nil {
		return nil, errors.New("la tienda no existe")
	}

	branchIds := a.BranchRepository.GetIdsByStoreID(createStoreCampaignDTO.StoreID)

	result := make(chan resultProcess)
	wg := &sync.WaitGroup{}
	wg.Add(len(branchIds))

	for _, branchID := range branchIds {
		go a.createBranchCampaign(createStoreCampaignDTO, branchID, result, wg)
	}

	go func() {
		wg.Wait()
		close(result)
	}()

	var response dto.StoreCampaignCreatedDTO

	for item := range result {
		if item.err != nil {
			err := dto.ErroBranchCampaign{
				Message:  item.err.Error(),
				BranchId: item.BranchID,
			}
			response.Errors = append(response.Errors, err)

			continue
		}

		response.BranchCampaigns = append(response.BranchCampaigns, *item.result)
	}

	return &response, nil
}

func (a *AddCampaignToStoreUC) createBranchCampaign(
	createStoreCampaignDTO dto.CreateStoreCampaignDTO,
	branchID uint,
	result chan resultProcess,
	wg *sync.WaitGroup,
) {
	defer wg.Done()

	branchCampaignDTO := dto.CreateBranchCampaignDTO{
		BranchID:      branchID,
		CampaignID:    createStoreCampaignDTO.CampaignID,
		StartDate:     createStoreCampaignDTO.StartDate,
		EndDate:       createStoreCampaignDTO.EndDate,
		Operator:      createStoreCampaignDTO.Operator,
		OperatorValue: createStoreCampaignDTO.OperatorValue,
		MinAmount:     createStoreCampaignDTO.MinAmount,
	}

	branchCampaign, err := a.CreateBranchCampaignUC.Execute(branchCampaignDTO)

	result <- resultProcess{
		result:   branchCampaign,
		err:      err,
		BranchID: branchID,
	}
}
