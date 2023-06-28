package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"lucio.com/order-service/src/domain/contracts/usecases"
	"lucio.com/order-service/src/domain/dto"
)

type StoreController struct {
	CreateStoreUC usecases.CreateStoreUC
}

func (s *StoreController) Create(ctx *gin.Context) {
	var createStoreDTO dto.CreateStoreDTO

	if err := ctx.ShouldBindJSON(&createStoreDTO); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
			"id":    "bad_request",
		})
		return
	}

	fmt.Print(createStoreDTO)

	store, err := s.CreateStoreUC.Execute(createStoreDTO)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
			"id":    "create_store_error",
		})
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"data": store,
	})
}
