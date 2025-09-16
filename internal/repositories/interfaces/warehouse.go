package interfaces

import (
	"context"
	"dropx/internal/domain/models"
	"dropx/pkg/utils"
	"github.com/gofrs/uuid"
)

type Warehouse interface {
	Create(ctx context.Context, warehouse *models.Warehouse) error
	FindByID(ctx context.Context, id uuid.UUID) (*models.Warehouse, error)
	Update(ctx context.Context, id uuid.UUID, fields map[string]interface{}) (*models.Warehouse, error)
	Delete(ctx context.Context, id uuid.UUID) error
	List(ctx context.Context, params utils.QueryParams) ([]models.Warehouse, int64, error)
}
