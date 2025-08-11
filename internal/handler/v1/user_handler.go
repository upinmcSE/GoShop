package v1handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	v1dto "github.com/upinmcSE/goshop/internal/dto/v1"
	v1service "github.com/upinmcSE/goshop/internal/service/v1"
	"github.com/upinmcSE/goshop/internal/utils"
	"github.com/upinmcSE/goshop/internal/validation"
)

type UserHandler struct {
	service v1service.UserService
}

func NewUserHandler(service v1service.UserService) *UserHandler {
	return &UserHandler{
		service: service,
	}
}

type GetUserByUuidParam struct {
	Uuid string `uri:"uuid" binding:"uuid"`
}

type GetUsersParams struct {
	Search string `form:"search" binding:"omitempty,min=3,max=50,search"`
	Page   int    `form:"page" binding:"omitempty,gte=1,lte=100"`
	Limit  int    `form:"limit" binding:"omitempty,gte=1,lte=100"`
}

func (uh *UserHandler) GetAllUsers(ctx *gin.Context) {
	var params GetUsersParams
	if err := ctx.ShouldBindQuery(&params); err != nil {
		utils.ResponseValidator(ctx, validation.HandleValidationErrors(err))
		return
	}

	if params.Page == 0 {
		params.Page = 1
	}

	if params.Limit == 0 {
		params.Limit = 10
	}


	utils.ResponseSuccess(ctx, http.StatusOK, "")
}

func (uh *UserHandler) CreateUser(ctx *gin.Context) {
	var input v1dto.CreateUserInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		utils.ResponseValidator(ctx, validation.HandleValidationErrors(err))
		return
	}

	

	utils.ResponseSuccess(ctx, http.StatusCreated, "")
}

func (uh *UserHandler) GetUserByUUID(ctx *gin.Context) {
	var params GetUserByUuidParam
	if err := ctx.ShouldBindUri(&params); err != nil {
		utils.ResponseValidator(ctx, validation.HandleValidationErrors(err))
		return
	}

	

	utils.ResponseSuccess(ctx, http.StatusOK, "")
}

func (uh *UserHandler) UpdateUser(ctx *gin.Context) {
	var params GetUserByUuidParam
	if err := ctx.ShouldBindUri(&params); err != nil {
		utils.ResponseValidator(ctx, validation.HandleValidationErrors(err))
		return
	}

	var input v1dto.UpdateUserInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		utils.ResponseValidator(ctx, validation.HandleValidationErrors(err))
		return
	}

	

	utils.ResponseSuccess(ctx, http.StatusOK, "")
}

func (uh *UserHandler) DeleteUser(ctx *gin.Context) {
	var params GetUserByUuidParam
	if err := ctx.ShouldBindUri(&params); err != nil {
		utils.ResponseValidator(ctx, validation.HandleValidationErrors(err))
		return
	}



	utils.ResponseStatusCode(ctx, http.StatusNoContent)
}
