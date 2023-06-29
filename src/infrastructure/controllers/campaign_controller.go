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

func (s *CampaignController) Create(ctx *gin.Context) {
	var createCampaignDTO dto.CreateCampaignDTO

	if err := ctx.ShouldBindJSON(&createCampaignDTO); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
			"id":    "bad_request",
		})
		return
	}

	campaignCreated, err := s.CreateCampaignUC.Execute(createCampaignDTO)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
			"id":    "create_store_error",
		})
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"data": campaignCreated,
	})
}

func (s *CampaignController) AddCampaignToBranch(ctx *gin.Context) {

}
