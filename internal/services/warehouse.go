package services

import (
	"context"
	"dropx/internal/api/dto/request"
	"dropx/internal/domain/models"
	"dropx/internal/repositories"
	"dropx/internal/repositories/interfaces"
	"dropx/pkg/utils"
	"github.com/gofrs/uuid"
)

type WarehouseService struct {
	repo interfaces.Warehouse
}

func NewWarehouseService(repo repositories.Repositories) *WarehouseService {
	return &WarehouseService{repo: repo.WarehouseRepository}
}

func (w *WarehouseService) Create(
	ctx context.Context,
	req request.CreateWarehouseRequest,
) (*models.Warehouse, error) {
	warehouse := models.Warehouse{
		ID:       utils.GetId(),
		Name:     req.Name,
		Address:  req.Address,
		Capacity: req.Capacity,
		Status:   req.Status,
	}
	err := w.repo.Create(ctx, &warehouse)
	if err != nil {
		return nil, err
	}

	return &warehouse, nil
}

func (w *WarehouseService) Get(
	ctx context.Context,
	id uuid.UUID,
) (*models.Warehouse, error) {
	return w.repo.FindByID(ctx, id)
}

func (w *WarehouseService) Update(
	ctx context.Context,
	id uuid.UUID,
	req request.UpdateWarehouselRequest,
) (*models.Warehouse, error) {
	return w.repo.Update(ctx, id, req.ToMap())
}

func (w *WarehouseService) Delete(
	ctx context.Context,
	id uuid.UUID,
) error {
	return w.repo.Delete(ctx, id)
}

func (w *WarehouseService) List(
	ctx context.Context,
	req request.ListRequest,
) ([]models.Warehouse, int64, error) {
	queryParam := utils.ParseQueryParams(
		req.Page,
		req.Size,
		req.Sort,
		req.Order,
		req.Filter,
	)
	return w.repo.List(ctx, queryParam)
}
