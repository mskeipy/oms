package models

import (
	"dropx/pkg/utils"
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type User struct {
	BaseModel
	ID       uuid.UUID `json:"id" gorm:"type:uuid;uniqueIndex;default:uuid_generate_v4()"`
	Name     string    `json:"name" gorm:"type:varchar(100);not null"`
	Email    string    `json:"email" gorm:"type:varchar(100);uniqueIndex;not null"`
	Password string    `json:"password" gorm:"type:varchar(255);not null"`
	Role     string    `json:"role" gorm:"type:varchar(50);not null;default:OA"`
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	hashedPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}
	u.Password = hashedPassword
	return nil
}
