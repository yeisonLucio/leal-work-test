package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"lucio.com/order-service/src/domain/contracts/usecases"
	"lucio.com/order-service/src/domain/dto"
)

type StoreController struct {
	CreateStoreUC usecases.CreateStoreUC
}

// @Summary Servicio para crear una tienda
// @Description Permite crear una determinada tienda
// @Tags Stores
// @Accept json
// @Produce json
// @Param body body dto.CreateStoreDTO true "Body data"
// @Success 201 {object} dto.StoreCreatedDTO
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /stores [post]
func (s *StoreController) Create(ctx *gin.Context) {
	var createStoreDTO dto.CreateStoreDTO

	if err := ctx.ShouldBindJSON(&createStoreDTO); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse{
			Message: err.Error(),
			ID:      "bad_request",
		})
		return
	}

	store, err := s.CreateStoreUC.Execute(createStoreDTO)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse{
			Message: err.Error(),
			ID:      "unexpected_error",
		})
	}

	ctx.JSON(http.StatusCreated, store)
}
