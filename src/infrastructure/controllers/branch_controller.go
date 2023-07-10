package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"lucio.com/order-service/src/domain/contracts/usecases"
	"lucio.com/order-service/src/domain/dto"
)

type BranchController struct {
	CreateBranchUC         usecases.CreateBranchUC
	CreateBranchCampaignUC usecases.CreateBranchCampaignUC
	AddCampaignToStoreUC   usecases.AddCampaignToStoreUC
	GetBranchCampaignsUC   usecases.GetBranchCampaignsUC
}

type errorResponse struct {
	ID      string `json:"id"`
	Message string `json:"message"`
}

// @Summary Servicio para crear sucursales de una tienda
// @Description Permite crear una determinada sucursal para una tienda
// @Tags Stores
// @Accept json
// @Produce json
// @Param store_id path int true "Store ID"
// @Param body body dto.CreateBranchDTO true "Body data"
// @Success 201 {object} dto.BranchCreatedDTO
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /stores/{store_id}/branches [post]
func (b *BranchController) Create(ctx *gin.Context) {
	var createBranchDTO dto.CreateBranchDTO

	if err := ctx.ShouldBindJSON(&createBranchDTO); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse{
			Message: err.Error(),
			ID:      "bad_request_error",
		})
		return
	}

	param := ctx.Param("store_id")

	storeID, err := strconv.Atoi(param)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse{
			Message: "el parámetro store_id debe ser un numero entero",
			ID:      "bad_request_error",
		})
		return
	}

	createBranchDTO.StoreID = uint(storeID)

	branchCreated, err := b.CreateBranchUC.Execute(createBranchDTO)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse{
			Message: err.Error(),
			ID:      "unexpected_error",
		})
		return
	}

	ctx.JSON(http.StatusCreated, branchCreated)
}

// @Summary Servicio para asociar una campaña a una sucursal
// @Description Permite asociar una campaña a una sucursal especifica
// @Tags Campaigns
// @Accept json
// @Produce json
// @Param campaign_id path int true "Campaign ID"
// @Param branch_id path int true  "Branch ID"
// @Param body body dto.CreateBranchCampaignDTO true "Body data"
// @Success 201 {object} dto.BranchCampaignCreatedDTO
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /campaigns/{campaign_id}/branches/{branch_id} [post]
func (b *BranchController) CreateBranchCampaign(ctx *gin.Context) {
	var createBranchCampaignDTO dto.CreateBranchCampaignDTO

	if err := ctx.ShouldBindJSON(&createBranchCampaignDTO); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse{
			Message: err.Error(),
			ID:      "bad_request",
		})
		return
	}

	paramCampaignID := ctx.Param("campaign_id")

	campaignID, err := strconv.Atoi(paramCampaignID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse{
			Message: "el parámetro campaign_id debe ser un un numero entero",
			ID:      "bad_request_error",
		})
		return
	}

	createBranchCampaignDTO.CampaignID = uint(campaignID)

	paramBranchID := ctx.Param("branch_id")

	branchID, err := strconv.Atoi(paramBranchID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse{
			Message: "el parámetro branch_id debe ser un un numero entero",
			ID:      "bad_request_error",
		})
		return
	}

	createBranchCampaignDTO.BranchID = uint(branchID)

	result, err := b.CreateBranchCampaignUC.Execute(createBranchCampaignDTO)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse{
			Message: err.Error(),
			ID:      "unexpected_error",
		})
		return
	}

	ctx.JSON(http.StatusCreated, result)
}

// @Summary Servicio para asociar una campaña a todas las sucursales de una tienda
// @Description Permite asociar una campaña a todas las sucursales de una tienda
// @Tags Campaigns
// @Accept json
// @Produce json
// @Param campaign_id path int true "Campaign ID"
// @Param store_id path int true "Store ID"
// @Param body body dto.CreateStoreCampaignDTO true "Body data"
// @Success 201 {object} dto.StoreCampaignCreatedDTO
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /campaigns/{campaign_id}/stores/{store_id} [post]
func (b *BranchController) AddCampaignToBranches(ctx *gin.Context) {
	var createStoreCampaignDTO dto.CreateStoreCampaignDTO

	if err := ctx.ShouldBindJSON(&createStoreCampaignDTO); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse{
			Message: err.Error(),
			ID:      "bad_request",
		})
		return
	}

	paramCampaignID := ctx.Param("campaign_id")

	campaignID, err := strconv.Atoi(paramCampaignID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse{
			Message: "el parámetro campaign_id debe ser un un numero entero",
			ID:      "bad_request_error",
		})
		return
	}

	createStoreCampaignDTO.CampaignID = uint(campaignID)

	paramStoreID := ctx.Param("store_id")

	storeID, err := strconv.Atoi(paramStoreID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse{
			Message: "el parámetro store_id debe ser un un numero entero",
			ID:      "bad_request_error",
		})
		return
	}

	createStoreCampaignDTO.StoreID = uint(storeID)

	result, err := b.AddCampaignToStoreUC.Execute(createStoreCampaignDTO)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse{
			Message: err.Error(),
			ID:      "unexpected_error",
		})
		return
	}

	ctx.JSON(http.StatusCreated, result)
}

// @Summary Servicio para obtener las campañas de una sucursal
// @Description Permite obtener las campañas de una determinada sucursal
// @Tags Campaigns
// @Accept json
// @Produce json
// @Param branch_id path int true "Branch ID"
// @Success 200 {object} dto.BranchCampaignReportDTO
// @Failure 400 {object} errorResponse
// @Router /campaigns/branches/{branch_id} [get]
func (b *BranchController) GetBranchCampaignsByBranch(ctx *gin.Context) {
	paramBranchID := ctx.Param("branch_id")

	branchID, err := strconv.Atoi(paramBranchID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse{
			Message: "el parámetro branch_id debe ser un un numero entero",
			ID:      "bad_request_error",
		})
		return
	}

	result := b.GetBranchCampaignsUC.Execute(uint(branchID))

	ctx.JSON(http.StatusOK, result)
}
