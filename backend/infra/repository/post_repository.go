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

func (r *PostRepository) Create(e *entity.Post) (uint32, error) {
	_, err := r.db.NewInsert().Model(e).Exec(context.Background())
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

func (r *PostRepository) GetRecentPosts(ctx context.Context, count int) ([]*entity.Post, error) {
	var posts []*entity.Post
	err := r.db.NewSelect().Model(&posts).Order("created_at DESC").Limit(10).Offset(count).Scan(ctx)
	if err != nil {
		return nil, err
	}
	return posts, nil
}
