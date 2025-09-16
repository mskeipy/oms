package interfaces

import (
	"dropx/internal/domain/models"
	"gorm.io/gorm"
)

type OrderItem interface {
	Create(
		warehouseOrderItems []*models.WarehouseOrderItem,
		tx *gorm.DB,
	) error
}
