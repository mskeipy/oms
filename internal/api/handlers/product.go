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

type ProductHandler struct {
	ProductService services.ProductService
}

func NewProductHandler(productService services.ProductService) *ProductHandler {
	return &ProductHandler{productService}
}

func (h *ProductHandler) Create(ctx *gin.Context) {
	var req request.CreateProductRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		utils.WriteErrorResponse(ctx, err, http.StatusBadRequest)
		return
	}

	product, err := h.ProductService.Create(ctx, req)
	if err != nil {
		utils.WriteErrorResponse(ctx, err, http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusCreated, response.CommonResponse{Data: product})
}

func (h *ProductHandler) Get(ctx *gin.Context) {
	id := ctx.Param(constants.Id)

	product, err := h.ProductService.Get(ctx, utils.StringToUUID(id))
	if err != nil {
		utils.WriteErrorResponse(ctx, err, http.StatusBadRequest)
		return
	}

	ctx.JSON(http.StatusOK, response.CommonResponse{
		Data: product,
	})
}

func (h *ProductHandler) Delete(ctx *gin.Context) {
	id := ctx.Param(constants.Id)

	err := h.ProductService.Delete(ctx, utils.StringToUUID(id))
	if err != nil {
		utils.WriteErrorResponse(ctx, err, http.StatusBadRequest)
		return
	}

	ctx.JSON(http.StatusOK, response.CommonResponse{
		Data: nil,
	})
}

func (h *ProductHandler) Update(ctx *gin.Context) {
	var req request.UpdateProductRequest
	id := ctx.Param(constants.Id)

	if err := ctx.ShouldBindJSON(&req); err != nil {
		utils.WriteErrorResponse(ctx, err, http.StatusBadRequest)
		return
	}

	product, err := h.ProductService.Update(ctx, utils.StringToUUID(id), req)
	if err != nil {
		utils.WriteErrorResponse(ctx, err, http.StatusBadRequest)
		return
	}

	ctx.JSON(http.StatusOK, response.CommonResponse{
		Data: product,
	})
}
func (h *ProductHandler) List(ctx *gin.Context) {
	var req request.ListRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		utils.WriteErrorResponse(ctx, err, http.StatusBadRequest)
		return
	}

	items, total, err := h.ProductService.List(ctx, req)
	if err != nil {
		utils.WriteErrorResponse(ctx, err, http.StatusBadRequest)
		return
	}

	res := response.PaginatedResponse{
		Data:       items,
		Pagination: response.NewPagination(req.Page, req.Size, total),
	}

	ctx.JSON(http.StatusOK, response.CommonResponse{
		Data: res,
	})
}

func (h *ProductHandler) CreateBundle(ctx *gin.Context) {
	var req request.CreateBundleRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		utils.WriteErrorResponse(ctx, err, http.StatusBadRequest)
		return
	}

	item, err := h.ProductService.CreateBundle(ctx.Request.Context(), req)
	if err != nil {
		utils.WriteErrorResponse(ctx, err, http.StatusBadRequest)
		return
	}

	ctx.JSON(http.StatusOK, response.CommonResponse{
		Data: item,
	})
}
