package router

import (
	"dropx/internal/api/handlers"
	"dropx/internal/api/middleware"
	"dropx/pkg/constants"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, handlers *handlers.AppHandlers) {
	api := r.Group("/api")

	auth := api.Group("/auth")
	auth.POST("/register", handlers.AuthHandler.Register)
	auth.POST("/login", handlers.AuthHandler.Login)
	auth.POST("/logout", handlers.AuthHandler.Logout)
	auth.POST("/forgot-password", handlers.AuthHandler.ForgotPassword)
	auth.POST("/reset-password", handlers.AuthHandler.ResetPassword)

	protected := api.Group("", middleware.AuthMiddleware())

	user := protected.Group(
		"/user",
		middleware.OnlyRoles(constants.RoleSA, constants.RoleOU),
	)
	user.GET("/:id", handlers.UserHandler.Get)
	user.GET("/me", handlers.UserHandler.GetMe)
	user.PUT("/update/:id", handlers.UserHandler.Update)

	warehouse := protected.Group(
		"/warehouse",
		middleware.OnlyRoles(constants.RoleSA, constants.RoleOU),
	)
	warehouse.POST("", handlers.WarehouseHandler.Create)
	warehouse.GET("/:id", handlers.WarehouseHandler.Get)
	warehouse.PUT("/:id", handlers.WarehouseHandler.Update)
	warehouse.DELETE("/:id", handlers.WarehouseHandler.Delete)
	warehouse.GET("", handlers.WarehouseHandler.List)

	product := protected.Group(
		"/product",
		middleware.OnlyRoles(constants.RoleSA, constants.RoleOU),
	)
	product.POST("", handlers.ProductHandler.Create)
	product.GET("/:id", handlers.ProductHandler.Get)
	product.GET("", handlers.ProductHandler.List)
	product.DELETE("/:id", handlers.ProductHandler.Delete)
	product.PUT("/:id", handlers.ProductHandler.Update)
	product.POST("/bundle", handlers.ProductHandler.CreateBundle)

	operation := protected.Group(
		"/operation",
		middleware.OnlyRoles(constants.RoleSA, constants.RoleOU),
	)
	operation.POST("/submit-warehouse-order", handlers.OperationHandler.SubmitWarehouseOrder)
}
