package services

import (
	"dropx/internal/repositories"
	"dropx/pkg/job/rabbitmq"
)

type AppServices struct {
	AuthService      *AuthServices
	UserService      *UserService
	WarehouseService *WarehouseService
	ProductService   *ProductService
	OperationService *OperationService
}

func NewAppServices(repo *repositories.Repositories, rabbit *rabbitmq.RabbitMQ) *AppServices {
	return &AppServices{
		AuthService:      NewAuthServices(*repo, rabbit),
		UserService:      NewUserService(*repo),
		WarehouseService: NewWarehouseService(*repo),
		ProductService:   NewProductService(*repo),
		OperationService: NewOperationService(*repo),
	}
}
