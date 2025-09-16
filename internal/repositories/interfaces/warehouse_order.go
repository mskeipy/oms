package interfaces

import (
	"dropx/internal/domain/models"
	"gorm.io/gorm"
)

type WarehouseOrder interface {
	Create(order *models.WarehouseOrder, tx *gorm.DB) error
}
