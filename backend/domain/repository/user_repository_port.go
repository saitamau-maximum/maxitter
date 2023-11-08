package repository

import (
	"context"

	"github.com/saitamau-maximum/maxitter/backend/domain/entity"
)

type IUserService interface {
	GetAll(ctx context.Context) (entity.UserSlice, error)
	GetByID(ctx context.Context, id string) (*entity.User, error)
	Create(ctx context.Context, user *entity.User) (*entity.User, error)
	Delete(ctx context.Context, id string) error
}
