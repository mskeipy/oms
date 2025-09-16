package models

import (
	"github.com/gofrs/uuid"
)

type Product struct {
	BaseModel
	ID          uuid.UUID       `json:"id" gorm:"type:uuid;uniqueIndex;default:uuid_generate_v4()"`
	SKU         string          `json:"sku" gorm:"type:varchar(255);unique;not null"`
	CustomerSKU string          `json:"customer_sku" gorm:"type:varchar(255)"`
	Name        string          `json:"name" gorm:"type:varchar(255);not null"`
	Description string          `json:"description" gorm:"type:text"`
	Dimension   string          `json:"dimension" gorm:"type:varchar(255)"`
	Weight      float64         `json:"weight" gorm:"type:decimal(10,2);default:0"`
	Status      string          `json:"status" gorm:"type:varchar(50);default:'pending'"`
	IsBundle    bool            `json:"is_bundle" gorm:"default:false"`
	BundleItems []ProductBundle `json:"bundle_items" gorm:"foreignKey:BundleID"`
}

type ProductBundle struct {
	ID               uuid.UUID `json:"id" gorm:"type:uuid;uniqueIndex;default:uuid_generate_v4()"`
	BundleID         uuid.UUID `json:"bundle_id" gorm:"type:uuid;not null"`
	ProductID        uuid.UUID `json:"product_id" gorm:"type:uuid;not null"`
	QuantityInBundle int32     `json:"quantity_in_bundle" gorm:"not null"`
	Product          Product   `json:"-" gorm:"foreignKey:ProductID;references:ID"`
}
