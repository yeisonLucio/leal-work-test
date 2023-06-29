package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"lucio.com/order-service/src/domain/contracts/usecases"
	"lucio.com/order-service/src/domain/dto"
)

type BranchController struct {
	CreateBranchUC usecases.CreateBranchUC
}

func (b *BranchController) Create(ctx *gin.Context) {
	var createBranchDTO dto.CreateBranchDTO

	if err := ctx.ShouldBindJSON(&createBranchDTO); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
			"id":    "bad_request_error",
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

	createBranchDTO.StoreID = uint(storeID)

	branchCreated, err := b.CreateBranchUC.Execute(createBranchDTO)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
			"id":    "unexpected_error",
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"data": branchCreated,
	})
}
