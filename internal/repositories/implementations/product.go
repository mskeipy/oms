package implementations

import (
	"context"
	"dropx/internal/domain/models"
	"dropx/internal/infrastructure/database"
	"dropx/internal/repositories/interfaces"
	"dropx/pkg/utils"
	"fmt"
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) interfaces.Product {
	return &productRepository{db: db}
}

func (repo productRepository) StandardQuery(ctx context.Context) *gorm.DB {
	return database.PostgresqlDB(ctx).Preload("BundleItems")
}

func (repo productRepository) Create(ctx context.Context, product *models.Product, tx *gorm.DB) error {
	return tx.Create(product).Error
}

func (repo productRepository) FindByID(ctx context.Context, id uuid.UUID) (*models.Product, error) {
	product := models.Product{ID: id}
	if err := repo.StandardQuery(ctx).First(&product).Error; err != nil {
		return nil, err
	}
	return &product, nil
}

func (repo productRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return database.Transaction(ctx, func(tx *gorm.DB) (err error) {
		product := models.Product{ID: id}
		return tx.Delete(&product).Error
	})
}

func (repo productRepository) Update(
	ctx context.Context,
	id uuid.UUID,
	fields map[string]interface{},
) (*models.Product, error) {
	var item models.Product
	err := database.Transaction(ctx,
		func(tx *gorm.DB) error {
			err := database.PostgresqlDB(ctx).Model(&models.Product{ID: id}).Updates(fields).Error
			if err != nil {
				return err
			}
			return tx.First(&item, "id = ?", id).Error
		},
	)
	return &item, err
}

func (repo productRepository) List(
	ctx context.Context,
	params utils.QueryParams,
) ([]models.Product, int64, error) {
	var items []models.Product
	var total int64

	query := repo.StandardQuery(ctx)
	query = utils.ApplyFilters(query, params.Filters)
	err := query.Order(params.SortBy + " " + params.Order).
		Limit(params.PageSize).
		Offset((params.Page - 1) * params.PageSize).
		Find(&items).
		Offset(-1).Limit(-1).Count(&total).
		Error
	if err != nil {
		return nil, 0, err
	}
	return items, total, err
}

func (repo productRepository) CreateBundleItems(ctx context.Context, items []models.ProductBundle, tx *gorm.DB) error {
	return tx.Create(&items).Error
}

func (repo productRepository) CheckExists(
	ctx context.Context,
	filters map[string]interface{},
) (bool, error) {
	var count int64
	var items []models.Product

	query := repo.StandardQuery(ctx)
	for field, value := range filters {
		query = query.Where(fmt.Sprintf("%s = ?", field), value)
	}

	if err := query.Find(&items).Count(&count).Error; err != nil {
		return false, err
	}

	return count > 0, nil
}
