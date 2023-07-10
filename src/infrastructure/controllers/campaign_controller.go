package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"lucio.com/order-service/src/domain/contracts/usecases"
	"lucio.com/order-service/src/domain/dto"
)

type CampaignController struct {
	CreateCampaignUC usecases.CreateCampaignUC
}

// @Summary Servicio para crear campañas
// @Description Permite crear una determinada campaña
// @Tags Campaigns
// @Accept json
// @Produce json
// @Param body body dto.CreateCampaignDTO true "Body data"
// @Success 201 {object} dto.CampaignCreatedDTO
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /campaigns [post]
func (s *CampaignController) Create(ctx *gin.Context) {
	var createCampaignDTO dto.CreateCampaignDTO

	if err := ctx.ShouldBindJSON(&createCampaignDTO); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse{
			Message: err.Error(),
			ID:      "bad_request",
		})
		return
	}

	campaignCreated, err := s.CreateCampaignUC.Execute(createCampaignDTO)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse{
			Message: err.Error(),
			ID:      "unexpected_error",
		})
	}

	ctx.JSON(http.StatusCreated, campaignCreated)
}
