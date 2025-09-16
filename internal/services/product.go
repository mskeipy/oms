package services

import (
	"context"
	"dropx/internal/api/dto/request"
	"dropx/internal/domain/models"
	"dropx/internal/infrastructure/database"
	"dropx/internal/repositories"
	"dropx/internal/repositories/interfaces"
	"dropx/pkg/constants"
	"dropx/pkg/utils"
	"errors"
	"fmt"
	"github.com/gofrs/uuid"
	"github.com/samber/lo"
	"gorm.io/gorm"
	"strings"
)

type ProductService struct {
	repo interfaces.Product
}

func NewProductService(repo repositories.Repositories) *ProductService {
	return &ProductService{repo: repo.ProductRepository}
}

func (s *ProductService) Create(
	ctx context.Context,
	req request.CreateProductRequest,
) (*models.Product, error) {
	isExist, err := s.repo.CheckExists(ctx, map[string]interface{}{
		"sku": req.SKU,
	})
	if err != nil {
		return nil, err
	}
	if isExist {
		return nil, errors.New("product sku already exists")
	}

	prod := models.Product{
		ID:          utils.GetId(),
		Name:        req.Name,
		SKU:         req.SKU,
		CustomerSKU: req.CustomerSKU,
		Description: req.Description,
		Dimension:   req.Dimension,
		Weight:      req.Weight,
	}
	err = database.Transaction(ctx, func(tx *gorm.DB) error {
		return s.repo.Create(ctx, &prod, tx)
	})
	if err != nil {
		return nil, err
	}

	return &prod, nil
}

func (s *ProductService) Get(
	ctx context.Context,
	id uuid.UUID,
) (*models.Product, error) {
	return s.repo.FindByID(ctx, id)
}

func (s *ProductService) Delete(
	ctx context.Context,
	id uuid.UUID,
) error {
	return s.repo.Delete(ctx, id)
}

func (s *ProductService) Update(
	ctx context.Context,
	id uuid.UUID,
	req request.UpdateProductRequest,
) (*models.Product, error) {
	return s.repo.Update(ctx, id, req.ToMap())
}

func (s *ProductService) List(
	ctx context.Context,
	req request.ListRequest,
) ([]models.Product, int64, error) {
	queryParam := utils.ParseQueryParams(
		req.Page,
		req.Size,
		req.Sort,
		req.Order,
		req.Filter,
	)
	return s.repo.List(ctx, queryParam)
}

func (s *ProductService) CreateBundle(
	ctx context.Context,
	req request.CreateBundleRequest,
) (*models.Product, error) {
	isExist, err := s.repo.CheckExists(ctx, map[string]interface{}{
		"sku": req.SKU,
	})
	if err != nil {
		return nil, err
	}
	if isExist {
		return nil, errors.New("product sku already exists")
	}

	itemIds := lo.Map(req.Items, func(item request.CreateBundleItem, _ int) string {
		return item.ProductID
	})

	queryParam := utils.ParseQueryParams(
		constants.DefaultPageNum,
		-1,
		constants.DefaultSort,
		constants.DefaultOrder,
		fmt.Sprintf("id: %s", strings.Join(itemIds, ",")),
	)
	_, itemTotal, err := s.repo.List(ctx, queryParam)
	if err != nil || itemTotal != int64(len(req.Items)) {
		return nil, errors.New("product not found")
	}

	bundle := &models.Product{
		SKU:         req.SKU,
		CustomerSKU: req.CustomerSKU,
		Name:        req.Name,
		Description: req.Description,
		Dimension:   req.Dimension,
		Weight:      req.Weight,
		IsBundle:    true,
	}

	var items []models.ProductBundle
	err = database.Transaction(ctx, func(tx *gorm.DB) error {
		if err = s.repo.Create(ctx, bundle, tx); err != nil {
			return err
		}
		for _, i := range req.Items {
			items = append(items, models.ProductBundle{
				ID:               utils.GetId(),
				BundleID:         bundle.ID,
				ProductID:        utils.StringToUUID(i.ProductID),
				QuantityInBundle: i.QuantityInBundle,
			})
		}
		return s.repo.CreateBundleItems(ctx, items, tx)
	})
	if err != nil {
		return nil, err
	}

	bundle.BundleItems = items
	return bundle, nil
}
