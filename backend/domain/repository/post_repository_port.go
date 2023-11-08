package repository

import (
	"context"

	"github.com/saitamau-maximum/maxitter/backend/domain/entity"
)

type IPostRepository interface {
	GetAll(ctx context.Context) (entity.PostSlice, error)
	GetByID(ctx context.Context, id string) (*entity.Post, error)
	GetRecent(ctx context.Context, count int) (entity.PostSlice, error)
	Create(ctx context.Context, post *entity.Post) (*entity.Post, error)
	Delete(ctx context.Context, id string) error
}
