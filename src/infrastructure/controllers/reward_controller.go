package controllers

import (
	"net/http"
	"strconv"

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

	param, found := ctx.Params.Get("store_id")
	if !found {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "el parámetro store_id es requerido",
			"id":    "bad_request_error",
		})
		return
	}

	storeID, err := strconv.Atoi(param)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "el parámetro store_id debe ser un un numero entero",
			"id":    "bad_request_error",
		})
		return
	}

	createRewardDTO.StoreID = uint(storeID)

	rewardCreated, err := s.CreateRewardUC.Execute(createRewardDTO)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
			"id":    "unexpected_error",
		})
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"data": rewardCreated,
	})
}
