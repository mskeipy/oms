package request

import "dropx/pkg/utils"

type UpdateUserRequest struct {
	Name     *string `json:"name" form:"name"`
	Password *string `json:"password" form:"password"`
	Role     *string `json:"role" form:"role"`
}

func (u UpdateUserRequest) ToMap() map[string]interface{} {
	var mapItem = map[string]interface{}{}
	if u.Name != nil {
		mapItem["name"] = *u.Name
	}
	if u.Password != nil {
		hashPassword, _ := utils.HashPassword(*u.Password)
		mapItem["password"] = hashPassword
	}
	if u.Role != nil {
		mapItem["role"] = *u.Name
	}
	return mapItem
}
