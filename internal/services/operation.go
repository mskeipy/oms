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
	"github.com/samber/lo"
	"gorm.io/gorm"
	"strings"
)

type OperationService struct {
	warehouseRepo      interfaces.Warehouse
	productRepo        interfaces.Product
	userRepo           interfaces.User
	warehouseOrderRepo interfaces.WarehouseOrder
	warehouseOrderItem interfaces.OrderItem
}

func NewOperationService(repo repositories.Repositories) *OperationService {
	return &OperationService{
		warehouseRepo:      repo.WarehouseRepository,
		productRepo:        repo.ProductRepository,
		userRepo:           repo.UserRepository,
		warehouseOrderRepo: repo.WarehouseOrderRepository,
		warehouseOrderItem: repo.OrderItemRepository,
	}
}

func (s *OperationService) SubmitOrder(
	ctx context.Context,
	req request.SubmitWarehouseOrderRequest,
) (*models.WarehouseOrder, error) {
	_, err := s.warehouseRepo.FindByID(ctx, utils.StringToUUID(*req.WarehouseID))
	if err != nil {
		return nil, fmt.Errorf("warehouse error: %w", err)
	}
	createdBy := ctx.Value(constants.TokenUserID).(string)
	if req.CreatedBy != nil {
		_, err = s.userRepo.FindByID(ctx, utils.StringToUUID(*req.CreatedBy))
		if err != nil {
			return nil, fmt.Errorf("craeted by error: %w", err)
		}
		createdBy = *req.CreatedBy
	}

	productIds := lo.Map(req.Items, func(item *request.SubmitWarehouseOrderItem, index int) string {
		return *item.ProductID
	})
	queryParam := utils.ParseQueryParams(
		constants.DefaultPageNum,
		-1,
		constants.DefaultSort,
		constants.DefaultOrder,
		fmt.Sprintf("id: %s", strings.Join(productIds, ",")),
	)
	_, itemTotal, err := s.productRepo.List(ctx, queryParam)
	if err != nil || itemTotal != int64(len(req.Items)) {
		return nil, errors.New("product not found")
	}

	var items []*models.WarehouseOrderItem

	order := models.WarehouseOrder{
		ID:          utils.GetId(),
		WarehouseID: utils.StringToUUID(*req.WarehouseID),
		OrderCode:   utils.GenerateOrderCode(constants.WarehouseInbound),
		Type:        constants.WarehouseInbound,
		Status:      constants.WarehouseOrderStatusDraft,
		CreatedBy:   utils.StringToUUID(createdBy),
	}

	for _, item := range req.Items {
		productId := utils.StringToUUID(*item.ProductID)
		expired := utils.StringToTime(*item.ExpiryDate)

		items = append(items, &models.WarehouseOrderItem{
			ID:         utils.GetId(),
			OrderID:    order.ID,
			ProductID:  productId,
			Quantity:   *item.Quantity,
			LotNo:      *item.LotNo,
			ExpiryDate: &expired,
		})
	}
	err = database.Transaction(ctx, func(tx *gorm.DB) error {
		err := s.warehouseOrderRepo.Create(&order, tx)
		if err != nil {
			return err
		}
		return s.warehouseOrderItem.Create(items, tx)
	})
	if err != nil {
		return nil, err
	}
	return &order, nil
}
