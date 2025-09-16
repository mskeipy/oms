package handlers

import (
	"dropx/internal/api/dto/request"
	"dropx/internal/api/dto/response"
	"dropx/internal/services"
	"dropx/pkg/constants"
	"dropx/pkg/utils"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"net/http"
)

type UserHandler struct {
	UserService services.UserService
}

func NewUserHandler(userService services.UserService) *UserHandler {
	return &UserHandler{UserService: userService}
}

func (u *UserHandler) Get(ctx *gin.Context) {
	id := ctx.Param(constants.Id)

	user, err := u.UserService.Get(ctx, utils.StringToUUID(id))
	if err != nil {
		utils.WriteErrorResponse(ctx, err, http.StatusBadRequest)
		return
	}

	ctx.JSON(http.StatusOK, response.CommonResponse{
		Data: user,
	})
}

func (u *UserHandler) GetMe(ctx *gin.Context) {
	id, exists := ctx.Get(constants.TokenUserID)
	if !exists {
		utils.WriteErrorResponse(ctx, errors.New("unauthorized"), http.StatusUnauthorized)
		return
	}
	user, err := u.UserService.Get(ctx, utils.StringToUUID(id.(string)))
	if err != nil {
		utils.WriteErrorResponse(ctx, err, http.StatusBadRequest)
		return
	}

	ctx.JSON(http.StatusOK, response.CommonResponse{
		Data: user,
	})
}

func (u *UserHandler) Update(ctx *gin.Context) {
	var req request.UpdateUserRequest
	id := ctx.Param(constants.Id)

	if err := ctx.ShouldBindJSON(&req); err != nil {
		utils.WriteErrorResponse(ctx, err, http.StatusBadRequest)
		return
	}

	user, err := u.UserService.Update(ctx, utils.StringToUUID(id), req)
	if err != nil {
		utils.WriteErrorResponse(ctx, err, http.StatusBadRequest)
		return
	}

	ctx.JSON(http.StatusOK, response.CommonResponse{
		Data: user,
	})
}
