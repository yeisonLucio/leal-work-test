package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"lucio.com/order-service/src/domain/contracts/usecases"
	"lucio.com/order-service/src/domain/dto"
)

type RewardController struct {
	CreateRewardUC usecases.CreateRewardUC
}

func (s *RewardController) Create(ctx *gin.Context) {
	var createRewardDTO dto.CreateRewardDTO

	if err := ctx.ShouldBindJSON(&createRewardDTO); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
			"id":    "bad_request",
		})
		return
	}

	rewardCreated, err := s.CreateRewardUC.Execute(createRewardDTO)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
			"id":    "create_store_error",
		})
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"data": rewardCreated,
	})
}
