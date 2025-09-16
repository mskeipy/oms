package implementations

import (
	"context"
	"dropx/internal/domain/models"
	"dropx/internal/infrastructure/database"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"time"
)

type inventoryRepository struct {
	db *gorm.DB
}

func NewInventoryRepository(db *gorm.DB) *inventoryRepository {
	return &inventoryRepository{db: db}
}

func (repo inventoryRepository) StandardQuery(ctx context.Context) *gorm.DB {
	return database.PostgresqlDB(ctx)
}

func (repo inventoryRepository) UpsertInventory(
	items []*models.WarehouseInventory,
	tx *gorm.DB,
) error {
	return tx.Clauses(clause.OnConflict{
		Columns: []clause.Column{
			{Name: "warehouse_id"},
			{Name: "product_id"},
			{Name: "lot_no"},
			{Name: "expiry_date"},
		},
		DoUpdates: clause.Assignments(map[string]interface{}{
			"quantity":   gorm.Expr("warehouse_inventories.quantity + EXCLUDED.quantity"),
			"updated_at": time.Now(),
		}),
	}).Model(&models.WarehouseInventory{}).Create(items).Error
}
