package handlers

import (
	"dropx/internal/api/dto/request"
	"dropx/internal/api/dto/response"
	"dropx/internal/services"
	"dropx/pkg/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AuthHandler struct {
	AuthService services.AuthServices
}

func NewAuthHandler(authService services.AuthServices) *AuthHandler {
	return &AuthHandler{AuthService: authService}
}

func (h *AuthHandler) Register(ctx *gin.Context) {
	var req request.RegisterRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		utils.WriteErrorResponse(ctx, err, http.StatusBadRequest)
		return
	}

	user, err := h.AuthService.Register(ctx, req)
	if err != nil {
		utils.WriteErrorResponse(ctx, err, http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusCreated, response.CommonResponse{Data: user})
}

func (h *AuthHandler) Login(ctx *gin.Context) {
	var req request.LoginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		utils.WriteErrorResponse(ctx, err, http.StatusBadRequest)
		return
	}

	token, err := h.AuthService.Authenticate(ctx, req)
	if err != nil {
		utils.WriteErrorResponse(ctx, err, http.StatusUnauthorized)
		return
	}

	ctx.JSON(http.StatusOK, response.CommonResponse{
		Data: response.LoginResponse{Email: req.Email, Token: token},
	})
}

func (h *AuthHandler) Logout(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Logged out successfully",
	})
}

func (h *AuthHandler) ForgotPassword(ctx *gin.Context) {
	var req request.ForgotPasswordRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		utils.WriteErrorResponse(ctx, err, http.StatusBadRequest)
		return
	}

	err := h.AuthService.GenerateAndSendResetToken(ctx, req.Email)
	if err != nil {
		utils.WriteErrorResponse(ctx, err, http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Reset link sent to your email."})
}

func (h *AuthHandler) ResetPassword(ctx *gin.Context) {
	var req request.ResetPasswordRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		utils.WriteErrorResponse(ctx, err, http.StatusBadRequest)
		return
	}

	err := h.AuthService.ResetPassword(ctx.Request.Context(), req)
	if err != nil {
		utils.WriteErrorResponse(ctx, err, http.StatusUnauthorized)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Password has been reset"})
}
