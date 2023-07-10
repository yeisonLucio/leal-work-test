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

// @Summary Servicio para crear premios de una tienda
// @Description Permite crear un premio para un comercio
// @Tags Stores
// @Accept json
// @Produce json
// @Param store_id path int true "Store ID"
// @Param body body dto.CreateRewardDTO true "Body data"
// @Success 201 {object} dto.RewardCreatedDTO
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /stores/{store_id}/rewards [post]
func (s *RewardController) Create(ctx *gin.Context) {
	var createRewardDTO dto.CreateRewardDTO

	if err := ctx.ShouldBindJSON(&createRewardDTO); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse{
			Message: err.Error(),
			ID:      "bad_request",
		})
		return
	}

	param, found := ctx.Params.Get("store_id")
	if !found {
		ctx.JSON(http.StatusBadRequest, errorResponse{
			Message: "el parámetro store_id es requerido",
			ID:      "bad_request_error",
		})
		return
	}

	storeID, err := strconv.Atoi(param)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse{
			Message: "el parámetro store_id debe ser un un numero entero",
			ID:      "bad_request_error",
		})
		return
	}

	createRewardDTO.StoreID = uint(storeID)

	rewardCreated, err := s.CreateRewardUC.Execute(createRewardDTO)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse{
			Message: err.Error(),
			ID:      "unexpected_error",
		})
	}

	ctx.JSON(http.StatusCreated, rewardCreated)
}
