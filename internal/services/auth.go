package services

import (
	"context"
	"dropx/internal/api/dto/request"
	"dropx/internal/domain/models"
	"dropx/internal/repositories"
	"dropx/internal/repositories/interfaces"
	"dropx/pkg/config"
	"dropx/pkg/constants"
	"dropx/pkg/job/rabbitmq"
	"dropx/pkg/job/rabbitmq/consumer"
	"dropx/pkg/utils"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"html/template"
	"log"
	"strings"
)

type AuthServices struct {
	UserRepo interfaces.User
	rabbit   *rabbitmq.RabbitMQ
}

func NewAuthServices(repo repositories.Repositories, rabbit *rabbitmq.RabbitMQ) *AuthServices {
	return &AuthServices{
		UserRepo: repo.UserRepository,
		rabbit:   rabbit,
	}
}

func (s *AuthServices) Register(ctx context.Context, req request.RegisterRequest) (*models.User, error) {
	existingUser, _ := s.UserRepo.FindByEmail(ctx, req.Email)
	if existingUser != nil {
		return nil, errors.New("email already registered")
	}

	user := &models.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
		Role:     req.Role,
	}
	err := s.UserRepo.Create(ctx, user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *AuthServices) Authenticate(ctx context.Context, req request.LoginRequest) (token string, err error) {
	user, err := s.UserRepo.FindByEmail(ctx, req.Email)
	if err != nil {
		return "", errors.New("invalid email or password")
	}

	match := utils.ComparePassword(user.Password, req.Password)
	if !match {
		return "", errors.New("invalid email or password")
	}

	return utils.GenerateToken(user.ID.String(), user.Email, user.Role)
}

func (s *AuthServices) GenerateAndSendResetToken(ctx *gin.Context, email string) error {
	user, err := s.UserRepo.FindByEmail(ctx, email)
	if err != nil {
		return fmt.Errorf("user not found")
	}

	resetToken, err := utils.GenerateResetPasswordToken(email)
	if err != nil {
		return fmt.Errorf("failed to generate reset token")
	}
	resetLink := strings.NewReplacer(
		":url", config.Global.AppApiUrl,
		":token", resetToken,
	).Replace(constants.ResetPasswordUrl)
	payload := consumer.ResetPasswordPayload{
		Email:     email,
		Name:      user.Name,
		ResetLink: template.URL(resetLink),
	}
	log.Println(payload)
	return s.rabbit.Publish(constants.ResetPasswordQueue, payload)
}

func (s *AuthServices) ResetPassword(ctx context.Context, req request.ResetPasswordRequest) error {
	tokenClaims, err := utils.VerifyResetPasswordToken(req.Token)
	if err != nil {
		log.Printf("failed to verify token: %v", err)
		return err
	}

	email := tokenClaims.Email
	user, err := s.UserRepo.FindByEmail(ctx, email)
	if err != nil {
		return err
	}
	password, err := utils.HashPassword(req.NewPassword)
	if err != nil {
		return err
	}
	_, err = s.UserRepo.Update(ctx, user.ID, map[string]interface{}{"password": password})
	return err
}
