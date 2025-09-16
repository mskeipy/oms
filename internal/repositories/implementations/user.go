package implementations

import (
	"context"
	"dropx/internal/domain/models"
	"dropx/internal/infrastructure/database"
	"dropx/internal/repositories/interfaces"
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) interfaces.User {
	return &UserRepository{db: db}
}

func (repo UserRepository) StandardQuery() *gorm.DB {
	return repo.db
}

func (UserRepository) Create(ctx context.Context, user *models.User) error {
	return database.Transaction(ctx, func(tx *gorm.DB) error {
		return database.PostgresqlDB(ctx).Create(user).Error
	})
}

func (repo UserRepository) FindByEmail(ctx context.Context, email string) (*models.User, error) {
	var user models.User
	if err := repo.StandardQuery().Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (repo UserRepository) FindByID(ctx context.Context, id uuid.UUID) (*models.User, error) {
	user := models.User{ID: id}
	if err := repo.StandardQuery().First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (repo UserRepository) Update(
	ctx context.Context,
	id uuid.UUID,
	fields map[string]interface{},
) (*models.User, error) {
	var user models.User
	err := database.Transaction(ctx,
		func(tx *gorm.DB) error {
			err := database.PostgresqlDB(ctx).Model(&models.User{ID: id}).Updates(fields).Error
			if err != nil {
				return err
			}
			return tx.First(&user, "id = ?", id).Error
		},
	)
	return &user, err
}
