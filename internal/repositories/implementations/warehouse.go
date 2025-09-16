package implementations

import (
	"context"
	"dropx/internal/domain/models"
	"dropx/internal/infrastructure/database"
	"dropx/internal/repositories/interfaces"
	"dropx/pkg/utils"
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type warehouseRepository struct {
	db *gorm.DB
}

func NewWarehouseRepository(db *gorm.DB) interfaces.Warehouse {
	return &warehouseRepository{db: db}
}

func (repo warehouseRepository) StandardQuery(ctx context.Context) *gorm.DB {
	return database.PostgresqlDB(ctx)
}

func (repo warehouseRepository) Create(ctx context.Context, warehouse *models.Warehouse) error {
	return database.Transaction(ctx, func(tx *gorm.DB) error {
		return database.PostgresqlDB(ctx).Create(warehouse).Error
	})
}

func (repo warehouseRepository) FindByID(ctx context.Context, id uuid.UUID) (*models.Warehouse, error) {
	warehouse := models.Warehouse{ID: id}
	if err := repo.StandardQuery(ctx).First(&warehouse).Error; err != nil {
		return nil, err
	}
	return &warehouse, nil
}

func (repo warehouseRepository) Update(ctx context.Context, id uuid.UUID, fields map[string]interface{}) (*models.Warehouse, error) {
	var warehouse models.Warehouse
	err := database.Transaction(ctx,
		func(tx *gorm.DB) error {
			err := database.PostgresqlDB(ctx).Model(&models.Warehouse{ID: id}).Updates(fields).Error
			if err != nil {
				return err
			}
			return tx.First(&warehouse, "id = ?", id).Error
		},
	)
	return &warehouse, err
}

func (repo warehouseRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return database.Transaction(ctx, func(tx *gorm.DB) (err error) {
		warehouse := models.Warehouse{ID: id}
		return tx.Delete(&warehouse).Error
	})
}

func (repo *warehouseRepository) List(ctx context.Context, params utils.QueryParams) ([]models.Warehouse, int64, error) {
	var warehouses []models.Warehouse
	var total int64

	query := repo.StandardQuery(ctx)

	query = utils.ApplyFilters(query, params.Filters)

	err := query.Order(params.SortBy + " " + params.Order).
		Limit(params.PageSize).
		Offset((params.Page - 1) * params.PageSize).
		Find(&warehouses).
		Offset(-1).Limit(-1).Count(&total).
		Error

	return warehouses, total, err
}
