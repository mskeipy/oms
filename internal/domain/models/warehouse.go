package models

import (
	"github.com/gofrs/uuid"
	"time"
)

type Warehouse struct {
	BaseModel
	ID       uuid.UUID `json:"id" gorm:"type:uuid;uniqueIndex;default:uuid_generate_v4()"`
	Name     string    `json:"name" gorm:"type:varchar(100);not null"`
	Capacity int32     `json:"capacity" gorm:"type:integer;not null"`
	Address  string    `json:"address" gorm:"type:varchar(255);not null"`
	Status   string    `json:"status" gorm:"type:varchar(50);not null"`
}

type WarehouseOrder struct {
	ID                uuid.UUID            `json:"id" gorm:"type:uuid;uniqueIndex;default:uuid_generate_v4()"`
	WarehouseID       uuid.UUID            `json:"warehouse_id" gorm:"type:uuid;not null"`
	OrderCode         string               `json:"order_code" gorm:"uniqueIndex;not null"`
	Type              string               `json:"type" gorm:"not null"`
	Status            string               `json:"status" gorm:"default:'draft'"`
	ContainerTracking string               `json:"container_tracking"`
	TrackingNumber    string               `json:"tracking_number"`
	ETA               *time.Time           `json:"eta"`
	RMACode           string               `json:"rma_code" gorm:"type:varchar(50)"`
	UrgentLevel       string               `json:"urgent_level"`
	CreatedBy         uuid.UUID            `json:"created_by" gorm:"type:uuid;not null"`
	OrderItems        []WarehouseOrderItem `gorm:"foreignKey:OrderID"`
	BaseModel
}

type WarehouseOrderItem struct {
	ID         uuid.UUID      `json:"id" gorm:"type:uuid;uniqueIndex;default:uuid_generate_v4()"`
	OrderID    uuid.UUID      `json:"order_id" gorm:"type:uuid;not null"`
	Order      WarehouseOrder `gorm:"foreignKey:OrderID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	ProductID  uuid.UUID      `json:"product_id" gorm:"type:uuid;not null"`
	Product    Product        `gorm:"foreignKey:ProductID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
	Quantity   int32          `json:"quantity" gorm:"not null"`
	LotNo      string         `json:"lot_no"`
	ExpiryDate *time.Time     `json:"expiry_date"`
	BaseModel
}

type WarehouseInventory struct {
	ID                uuid.UUID  `json:"id" gorm:"type:uuid;uniqueIndex;default:uuid_generate_v4()"`
	WarehouseID       uuid.UUID  `json:"warehouse_id" gorm:"type:uuid;not null;index:idx_inventory_unique,unique"`
	Warehouse         Warehouse  `gorm:"foreignKey:WarehouseID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
	ProductID         uuid.UUID  `json:"product_id" gorm:"type:uuid;not null;index:idx_inventory_unique,unique"`
	Product           Product    `gorm:"foreignKey:ProductID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
	LotNo             string     `json:"lot_no" gorm:"type:varchar(50);not null;index:idx_inventory_unique,unique"`
	ExpiryDate        *time.Time `json:"expiry_date" gorm:"not null;index:idx_inventory_unique,unique"`
	Quantity          int32      `json:"quantity" gorm:"not null"`
	AvailableQuantity int32      `json:"available_quantity" gorm:"not null"`
	DamagedQuantity   int32      `json:"damaged_quantity" gorm:"default:0"`
	HoldQuantity      int32      `json:"hold_quantity" gorm:"default:0"`
	BaseModel
}
