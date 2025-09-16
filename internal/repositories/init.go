package repositories

import (
	"dropx/internal/repositories/implementations"
	"dropx/internal/repositories/interfaces"
	"gorm.io/gorm"
)

type Repositories struct {
	UserRepository           interfaces.User
	WarehouseRepository      interfaces.Warehouse
	ProductRepository        interfaces.Product
	WarehouseOrderRepository interfaces.WarehouseOrder
	InventoryRepository      interfaces.Inventory
	OrderItemRepository      interfaces.OrderItem
}

func NewRepositories(db *gorm.DB) *Repositories {
	return &Repositories{
		UserRepository:           implementations.NewUserRepository(db),
		WarehouseRepository:      implementations.NewWarehouseRepository(db),
		ProductRepository:        implementations.NewProductRepository(db),
		InventoryRepository:      implementations.NewInventoryRepository(db),
		WarehouseOrderRepository: implementations.NewWarehouseOrderRepository(db),
		OrderItemRepository:      implementations.NewOrderItemRepository(db),
	}
}
