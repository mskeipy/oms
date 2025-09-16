package interfaces

import (
	"dropx/internal/domain/models"
	"gorm.io/gorm"
)

type Inventory interface {
	UpsertInventory(inventory []*models.WarehouseInventory, tx *gorm.DB) error
}
