package handlers

import (
	"dropx/internal/services"
)

type AppHandlers struct {
	AuthHandler      *AuthHandler
	UserHandler      *UserHandler
	WarehouseHandler *WarehouseHandler
	ProductHandler   *ProductHandler
	OperationHandler *OperationHandler
}

func NewAppHandlers(services *services.AppServices) *AppHandlers {
	return &AppHandlers{
		AuthHandler:      NewAuthHandler(*services.AuthService),
		UserHandler:      NewUserHandler(*services.UserService),
		WarehouseHandler: NewWarehouseHandler(*services.WarehouseService),
		ProductHandler:   NewProductHandler(*services.ProductService),
		OperationHandler: NewOperationHandler(*services.OperationService),
	}
}
