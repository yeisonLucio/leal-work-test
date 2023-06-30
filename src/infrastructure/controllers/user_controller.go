package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"lucio.com/order-service/src/domain/contracts/usecases"
	"lucio.com/order-service/src/domain/dto"
)

type UserController struct {
	CreateUserUC        usecases.CreateUserUC
	CreateTransactionUC usecases.CreateTransactionUC
}

func (u *UserController) Create(ctx *gin.Context) {
	var createUserDTO dto.CreateUserDTO

	if err := ctx.ShouldBindJSON(&createUserDTO); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
			"id":    "bad_request",
		})
		return
	}

	user, err := u.CreateUserUC.Execute(createUserDTO)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
			"id":    "unexpected_error",
		})
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"data": user,
	})
}

func (u *UserController) RegisterTransaction(ctx *gin.Context) {
	var createTransactionDTO dto.CreateTransactionDTO

	if err := ctx.ShouldBindJSON(&createTransactionDTO); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
			"id":    "bad_request",
		})
		return
	}

	paramUserID := ctx.Param("user_id")

	userID, err := strconv.Atoi(paramUserID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "el parámetro user_id debe ser un un numero entero",
			"id":    "bad_request_error",
		})
		return
	}

	paramBranchID := ctx.Param("branch_id")

	branchID, err := strconv.Atoi(paramBranchID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "el parámetro branch_id debe ser un un numero entero",
			"id":    "bad_request_error",
		})
		return
	}

	createTransactionDTO.BranchID = uint(branchID)
	createTransactionDTO.UserID = uint(userID)

	transaction, err := u.CreateTransactionUC.Execute(createTransactionDTO)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
			"id":    "unexpected_error",
		})
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"data": transaction,
	})
}
