package implementations

import (
	"context"
	"dropx/internal/domain/models"
	"dropx/internal/infrastructure/database"
	"gorm.io/gorm"
)

type orderItemRepository struct {
	db *gorm.DB
}

func NewOrderItemRepository(db *gorm.DB) *orderItemRepository {
	return &orderItemRepository{db: db}
}

func (repo orderItemRepository) StandardQuery(ctx context.Context) *gorm.DB {
	return database.PostgresqlDB(ctx)
}

func (repo orderItemRepository) Create(
	warehouseOrderItems []*models.WarehouseOrderItem,
	tx *gorm.DB,
) error {
	return tx.Create(warehouseOrderItems).Error
}
