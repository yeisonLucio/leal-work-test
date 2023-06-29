package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"lucio.com/order-service/src/domain/contracts/usecases"
	"lucio.com/order-service/src/domain/dto"
)

type UserController struct {
	CreateUserUC usecases.CreateUserUC
}

func (s *UserController) Create(ctx *gin.Context) {
	var createUserDTO dto.CreateUserDTO

	if err := ctx.ShouldBindJSON(&createUserDTO); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
			"id":    "bad_request",
		})
		return
	}

	user, err := s.CreateUserUC.Execute(createUserDTO)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
			"id":    "create_store_error",
		})
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"data": user,
	})
}
