package repository

import (
	"context"

	"github.com/uptrace/bun"

	"github.com/saitamau-maximum/maxitter/backend/internal/entity"
)

type PostRepository struct {
	db *bun.DB
}

func NewPostRepository(db *bun.DB) *PostRepository {
	return &PostRepository{
		db,
	}
}

func (r *PostRepository) Create(ctx context.Context, e *entity.Post) (string, error) {
	_, err := r.db.NewInsert().Model(e).Exec(ctx)
	if err != nil {
		return e.ID, err
	}
	return e.ID, nil
}

func (r *PostRepository) Find(ctx context.Context, id int) (*entity.Post, error) {
	post := &entity.Post{}
	err := r.db.NewSelect().Model(post).Where("id = ?", id).Scan(ctx)
	if err != nil {
		return nil, err
	}
	return post, nil
}

func (r *PostRepository) GetRecentPosts(ctx context.Context, count, offset int) ([]*entity.Post, error) {
	var posts []*entity.Post
	err := r.db.NewSelect().Model(&posts).Order("created_at DESC").Limit(count).Offset(offset).Scan(ctx)
	if err != nil {
		return nil, err
	}
	return posts, nil
}
