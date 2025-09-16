package handlers

import (
	"dropx/internal/api/dto/request"
	"dropx/internal/api/dto/response"
	"dropx/internal/services"
	"dropx/pkg/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type OperationHandler struct {
	Operation services.OperationService
}

func NewOperationHandler(operation services.OperationService) *OperationHandler {
	return &OperationHandler{operation}
}

func (h *OperationHandler) SubmitWarehouseOrder(ctx *gin.Context) {
	var req request.SubmitWarehouseOrderRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		utils.WriteErrorResponse(ctx, err, http.StatusBadRequest)
		return
	}

	order, err := h.Operation.SubmitOrder(ctx, req)
	if err != nil {
		utils.WriteErrorResponse(ctx, err, http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusCreated, response.CommonResponse{Data: order})
}
