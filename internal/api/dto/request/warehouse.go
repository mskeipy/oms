package request

type CreateWarehouseRequest struct {
	Name     string `json:"name" binding:"required"`
	Capacity int32  `json:"capacity" binding:"required"`
	Address  string `json:"address" binding:"required"`
	Status   string `json:"status" binding:"required,oneof=draft awaiting received completed canceled"`
}

type UpdateWarehouselRequest struct {
	Name     *string `json:"name" form:"name"`
	Capacity *int32  `json:"capacity" form:"capacity"`
	Address  *string `json:"address" form:"address"`
	Status   *string `json:"status" form:"status" binding:"omitempty,oneof=draft awaiting received completed canceled"`
}

func (w UpdateWarehouselRequest) ToMap() map[string]interface{} {
	var mapItem = map[string]interface{}{}
	if w.Name != nil {
		mapItem["name"] = *w.Name
	}
	if w.Capacity != nil {
		mapItem["capacity"] = *w.Capacity
	}
	if w.Address != nil {
		mapItem["address"] = *w.Address
	}
	if w.Status != nil {
		mapItem["status"] = *w.Status
	}
	return mapItem
}
