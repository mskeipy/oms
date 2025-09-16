package request

type SubmitWarehouseOrderRequest struct {
	WarehouseID *string                     `json:"warehouse_id" binding:"required,uuid"`
	CreatedBy   *string                     `json:"created_by" binding:"omitempty,uuid"`
	Items       []*SubmitWarehouseOrderItem `json:"items" binding:"required"`
}

type SubmitWarehouseOrderItem struct {
	ProductID  *string `json:"product_id" binding:"required,uuid"`
	Quantity   *int32  `json:"quantity" binding:"required"`
	LotNo      *string `json:"lot_no"`
	ExpiryDate *string `json:"expiry_date"`
}
