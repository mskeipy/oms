package implementations

import (
	"context"
	"dropx/internal/domain/models"
	"dropx/internal/infrastructure/database"
	"gorm.io/gorm"
)

type warehouseOrderRepository struct {
	db *gorm.DB
}

func NewWarehouseOrderRepository(db *gorm.DB) *warehouseOrderRepository {
	return &warehouseOrderRepository{db: db}
}

func (repo warehouseOrderRepository) StandardQuery(ctx context.Context) *gorm.DB {
	return database.PostgresqlDB(ctx)
}

func (repo warehouseOrderRepository) Create(
	warehouseOrder *models.WarehouseOrder,
	tx *gorm.DB,
) error {
	return tx.Create(warehouseOrder).Error
}
