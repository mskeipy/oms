package interfaces

import (
	"context"
	"dropx/internal/domain/models"
	"github.com/gofrs/uuid"
)

type User interface {
	Create(ctx context.Context, user *models.User) error
	FindByEmail(ctx context.Context, email string) (*models.User, error)
	FindByID(ctx context.Context, id uuid.UUID) (*models.User, error)
	Update(ctx context.Context, id uuid.UUID, fields map[string]interface{}) (*models.User, error)
}
