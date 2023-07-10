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

// @Summary Servicio para crear un usuario
// @Description Permite crear una determinado usuario
// @Tags Users
// @Accept json
// @Produce json
// @Param body body dto.CreateUserDTO true "Body data"
// @Success 201 {object} dto.UserCreatedDTO
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /users [post]
func (u *UserController) Create(ctx *gin.Context) {
	var createUserDTO dto.CreateUserDTO

	if err := ctx.ShouldBindJSON(&createUserDTO); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse{
			Message: err.Error(),
			ID:      "bad_request",
		})
		return
	}

	user, err := u.CreateUserUC.Execute(createUserDTO)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse{
			Message: err.Error(),
			ID:      "unexpected_error",
		})
	}

	ctx.JSON(http.StatusCreated, user)
}

// @Summary Servicio para crear transacciones de un usuario
// @Description Permite Registrar una transacción de un usuario en una sucursal
// @Tags Users
// @Accept json
// @Produce json
// @param user_id path int true "User ID"
// @param branch_id path int true "Branch ID"
// @Param body body dto.CreateTransactionDTO true "Body data"
// @Success 201 {object} dto.TransactionCreatedDTO
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /users/{user_id}/transactions/branches/{branch_id} [post]
func (u *UserController) RegisterTransaction(ctx *gin.Context) {
	var createTransactionDTO dto.CreateTransactionDTO

	if err := ctx.ShouldBindJSON(&createTransactionDTO); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse{
			Message: err.Error(),
			ID:      "bad_request",
		})
		return
	}

	paramUserID := ctx.Param("user_id")

	userID, err := strconv.Atoi(paramUserID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse{
			Message: "el parámetro user_id debe ser un un numero entero",
			ID:      "bad_request_error",
		})
		return
	}

	paramBranchID := ctx.Param("branch_id")

	branchID, err := strconv.Atoi(paramBranchID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse{
			Message: "el parámetro branch_id debe ser un un numero entero",
			ID:      "bad_request_error",
		})
		return
	}

	createTransactionDTO.BranchID = uint(branchID)
	createTransactionDTO.UserID = uint(userID)

	transaction, err := u.CreateTransactionUC.Execute(createTransactionDTO)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse{
			Message: err.Error(),
			ID:      "unexpected_error",
		})
	}

	ctx.JSON(http.StatusCreated, transaction)
}
