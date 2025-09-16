package request

type CreateProductRequest struct {
	SKU         string  `json:"sku" binding:"required"`
	CustomerSKU string  `json:"customer_sku"`
	Name        string  `json:"name" binding:"required"`
	Description string  `json:"description"`
	Dimension   string  `json:"dimension" binding:"required"`
	Weight      float64 `json:"weight" binding:"required"`
}
type UpdateProductRequest struct {
	Name        *string  `json:"name" form:"name"`
	CustomerSKU *string  `json:"customer_sku" form:"customer_sku"`
	Description *string  `json:"description" form:"description"`
	Dimension   *string  `json:"dimension" form:"dimension"`
	Weight      *float64 `json:"weight" form:"weight"`
	Status      *string  `json:"status" form:"status" binding:"omitempty,oneof=approved denied pending"`
}

func (w UpdateProductRequest) ToMap() map[string]interface{} {
	var mapItem = map[string]interface{}{}
	if w.Name != nil {
		mapItem["name"] = *w.Name
	}
	if w.CustomerSKU != nil {
		mapItem["customer_sku"] = *w.CustomerSKU
	}
	if w.Description != nil {
		mapItem["description"] = *w.Description
	}
	if w.Dimension != nil {
		mapItem["dimension"] = *w.Dimension
	}
	if w.Weight != nil {
		mapItem["weight"] = *w.Weight
	}
	if w.Status != nil {
		mapItem["status"] = *w.Status
	}
	return mapItem
}

type CreateBundleItem struct {
	ProductID        string `json:"product_id" binding:"required,uuid"`
	QuantityInBundle int32  `json:"quantity_in_bundle" binding:"required,min=1"`
}

type CreateBundleRequest struct {
	CreateProductRequest
	Items []CreateBundleItem `json:"items" binding:"required,dive"`
}
