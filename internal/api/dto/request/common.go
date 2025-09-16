package request

type ListRequest struct {
	Page   int    `json:"page" form:"page,default=1"`
	Size   int    `json:"size" form:"size,default=10"`
	Sort   string `json:"sort" form:"sort,default=created_at"`
	Order  string `json:"order" form:"order,default=desc" binding:"oneof=desc asc"`
	Filter string `json:"filter" form:"filter"`
}
