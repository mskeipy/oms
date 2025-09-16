package utils

import (
	"dropx/pkg/constants"
	"gorm.io/gorm"
	"strings"
)

type QueryParams struct {
	Page     int                    `json:"page"`
	PageSize int                    `json:"page_size"`
	SortBy   string                 `json:"sort_by"`
	Order    string                 `json:"order"` // asc, desc
	Filters  map[string]interface{} `json:"filters"`
}

func ParseQueryParams(page, size int, sort, order string, filter string) QueryParams {
	if page <= 0 {
		page = constants.DefaultPageNum
	}
	if size <= 0 {
		size = constants.DefaultPageSize
	}
	if sort == "" {
		sort = constants.DefaultSort
	}
	if order == "" {
		order = constants.DefaultOrder
	}

	return QueryParams{
		Page:     page,
		PageSize: size,
		SortBy:   sort,
		Order:    order,
		Filters:  ParseFilterString(filter),
	}
}

func ParseFilterString(filterStr string) map[string]interface{} {
	filters := make(map[string]interface{})
	if filterStr == "" {
		return filters
	}

	pairs := strings.Split(filterStr, ";")
	for _, pair := range pairs {
		kv := strings.SplitN(pair, ":", 2)
		if len(kv) == 2 {
			key := strings.TrimSpace(kv[0])
			value := strings.TrimSpace(kv[1])
			if strings.Contains(value, ",") {
				filters[key] = strings.Split(value, ",")
			} else {
				filters[key] = value
			}
		}
	}
	return filters
}

func ApplyFilters(db *gorm.DB, filters map[string]interface{}) *gorm.DB {
	for key, value := range filters {
		switch v := value.(type) {
		case string:
			db = db.Where(key+" = ?", v)
		case int, float64:
			db = db.Where(key+" = ?", v)
		case []string:
			db = db.Where(key+" IN ?", v)
		case []interface{}:
			db = db.Where(key+" IN ?", v)
		}
	}
	return db
}
