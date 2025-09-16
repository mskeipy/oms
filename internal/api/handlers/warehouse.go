package handlers

import (
	"dropx/internal/api/dto/request"
	"dropx/internal/api/dto/response"
	"dropx/internal/services"
	"dropx/pkg/constants"
	"dropx/pkg/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type WarehouseHandler struct {
	WarehouseService services.WarehouseService
}

func NewWarehouseHandler(warehouseService services.WarehouseService) *WarehouseHandler {
	return &WarehouseHandler{warehouseService}
}

func (h *WarehouseHandler) Create(ctx *gin.Context) {
	var req request.CreateWarehouseRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		utils.WriteErrorResponse(ctx, err, http.StatusBadRequest)
		return
	}

	warehouse, err := h.WarehouseService.Create(ctx, req)
	if err != nil {
		utils.WriteErrorResponse(ctx, err, http.StatusBadRequest)
		return
	}

	ctx.JSON(http.StatusCreated, response.CommonResponse{Data: warehouse})
}

func (h *WarehouseHandler) Get(ctx *gin.Context) {
	id := ctx.Param(constants.Id)

	warehouse, err := h.WarehouseService.Get(ctx, utils.StringToUUID(id))
	if err != nil {
		utils.WriteErrorResponse(ctx, err, http.StatusBadRequest)
		return
	}

	ctx.JSON(http.StatusOK, response.CommonResponse{
		Data: warehouse,
	})
}

func (h *WarehouseHandler) Update(ctx *gin.Context) {
	var req request.UpdateWarehouselRequest
	id := ctx.Param(constants.Id)

	if err := ctx.ShouldBindJSON(&req); err != nil {
		utils.WriteErrorResponse(ctx, err, http.StatusBadRequest)
		return
	}

	warehouse, err := h.WarehouseService.Update(ctx, utils.StringToUUID(id), req)
	if err != nil {
		utils.WriteErrorResponse(ctx, err, http.StatusBadRequest)
		return
	}

	ctx.JSON(http.StatusOK, response.CommonResponse{
		Data: warehouse,
	})
}

func (h *WarehouseHandler) Delete(ctx *gin.Context) {
	id := ctx.Param(constants.Id)

	err := h.WarehouseService.Delete(ctx, utils.StringToUUID(id))
	if err != nil {
		utils.WriteErrorResponse(ctx, err, http.StatusBadRequest)
		return
	}

	ctx.JSON(http.StatusOK, response.CommonResponse{
		Data: nil,
	})
}

func (h *WarehouseHandler) List(ctx *gin.Context) {
	var req request.ListRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		utils.WriteErrorResponse(ctx, err, http.StatusBadRequest)
		return
	}

	warehouses, total, err := h.WarehouseService.List(ctx, req)
	if err != nil {
		utils.WriteErrorResponse(ctx, err, http.StatusBadRequest)
		return
	}

	res := response.PaginatedResponse{
		Data:       warehouses,
		Pagination: response.NewPagination(req.Page, req.Size, total),
	}

	ctx.JSON(http.StatusOK, response.CommonResponse{
		Data: res,
	})
}
