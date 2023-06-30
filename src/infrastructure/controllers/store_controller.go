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
// @Success 200 {object} dto.StoreCreatedDTO
// @Router /stores [post]
func (s *StoreController) Create(ctx *gin.Context) {
	var createStoreDTO dto.CreateStoreDTO

	if err := ctx.ShouldBindJSON(&createStoreDTO); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
			"id":    "bad_request",
		})
		return
	}

	store, err := s.CreateStoreUC.Execute(createStoreDTO)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
			"id":    "unexpected_error",
		})
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"data": store,
	})
}
