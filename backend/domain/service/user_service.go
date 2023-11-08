package service

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/saitamau-maximum/maxitter/backend/domain/entity"
	"github.com/saitamau-maximum/maxitter/backend/domain/repository"
)

type IUserService interface {
	GetUsers(ctx context.Context) (entity.UserSlice, error)
	CreateUser(ctx context.Context, user *entity.User) (*entity.User, error)
}

type UserService struct {
	UserRepository repository.IUserService
}

func NewUserService(repo repository.IUserService) *UserService {
	return &UserService{UserRepository: repo}
}

func (s *UserService) GetUsers(ctx context.Context) (entity.UserSlice, error) {
	users, err := s.UserRepository.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (s *UserService) CreateUser(ctx context.Context, user *entity.User) (*entity.User, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	user.ID = id.String()
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	_, err = s.UserRepository.Create(ctx, user)
	if err != nil {
		return nil, err
	}

	return user, nil
}
