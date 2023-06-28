package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"lucio.com/order-service/src/domain/contracts/usecases"
	"lucio.com/order-service/src/domain/dto"
)

type BranchController struct {
	CreateBranchUC usecases.CreateBranchUC
}

func (b *BranchController) Create(ctx *gin.Context) {
	var CreateBranchDTO dto.CreateBranchDTO

	if err := ctx.ShouldBindJSON(&CreateBranchDTO); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
			"id":    "bad_request_error",
		})

		return
	}

	branchCreated, err := b.CreateBranchUC.Execute(CreateBranchDTO)
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
