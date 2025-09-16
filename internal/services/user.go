package services

import (
	"context"
	"dropx/internal/api/dto/request"
	"dropx/internal/domain/models"
	"dropx/internal/repositories"
	"dropx/internal/repositories/interfaces"
	"github.com/gofrs/uuid"
)

type UserService struct {
	UserRepo interfaces.User
}

func NewUserService(repo repositories.Repositories) *UserService {
	return &UserService{UserRepo: repo.UserRepository}
}

func (s *UserService) Get(ctx context.Context, id uuid.UUID) (*models.User, error) {
	return s.UserRepo.FindByID(ctx, id)
}

func (s *UserService) Update(ctx context.Context, id uuid.UUID, req request.UpdateUserRequest) (*models.User, error) {
	return s.UserRepo.Update(ctx, id, req.ToMap())
}
