package interfaces

import (
	"context"
	"dropx/internal/domain/models"
	"dropx/pkg/utils"
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type Product interface {
	Create(ctx context.Context, product *models.Product, tx *gorm.DB) error
	FindByID(ctx context.Context, id uuid.UUID) (*models.Product, error)
	Delete(ctx context.Context, id uuid.UUID) error
	Update(ctx context.Context, id uuid.UUID, fields map[string]interface{}) (*models.Product, error)
	List(ctx context.Context, params utils.QueryParams) ([]models.Product, int64, error)
	CreateBundleItems(ctx context.Context, items []models.ProductBundle, tx *gorm.DB) error
	CheckExists(ctx context.Context, fields map[string]interface{}) (bool, error)
}
