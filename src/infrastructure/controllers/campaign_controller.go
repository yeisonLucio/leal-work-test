package controllers

import (
	"net/http"
	"strconv"

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

func (s *CampaignController) AddCampaignToStore(ctx *gin.Context) {
	var createBranchCampaignDTO dto.CreateBranchCampaignDTO

	if err := ctx.ShouldBindJSON(&createBranchCampaignDTO); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
			"id":    "bad_request",
		})
		return
	}

	paramCampaignID := ctx.Param("campaign_id")

	campaignID, err := strconv.Atoi(paramCampaignID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "el parámetro campaign_id debe ser un un numero entero",
			"id":    "bad_request_error",
		})
		return
	}

	createBranchCampaignDTO.CampaignID = uint(campaignID)

	paramBranchID := ctx.Param("branch_id")

	branchID, err := strconv.Atoi(paramBranchID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "el parámetro branch_id debe ser un un numero entero",
			"id":    "bad_request_error",
		})
		return
	}

	createBranchCampaignDTO.BranchID = uint(branchID)
}
