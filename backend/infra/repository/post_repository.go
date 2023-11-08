package repository

import (
	"context"

	"github.com/saitamau-maximum/maxitter/backend/domain/entity"
	"github.com/saitamau-maximum/maxitter/backend/domain/repository"
	"github.com/saitamau-maximum/maxitter/backend/usecase/model"
	"github.com/uptrace/bun"
)

type PostRepository struct {
	DB *bun.DB
}

func NewPostRepository(db *bun.DB) repository.IPostRepository {
	return &PostRepository{DB: db}
}

func (r *PostRepository) GetAll(ctx context.Context) (entity.PostSlice, error) {
	var modelPosts []model.Post
	err := r.DB.NewSelect().Model(&modelPosts).Scan(ctx)
	if err != nil {
		return nil, err
	}

	var posts entity.PostSlice
	for _, p := range modelPosts {
		posts = append(posts, p.ToPostEntity())
	}
	return posts, nil
}

func (r *PostRepository) GetByID(ctx context.Context, id string) (*entity.Post, error) {
	modelPosts := &model.Post{}
	err := r.DB.NewSelect().Model(modelPosts).Where("id = ?", id).Scan(ctx)
	if err != nil {
		return nil, err
	}
	return modelPosts.ToPostEntity(), nil
}

func (r *PostRepository) GetRecent(ctx context.Context, count int) (entity.PostSlice, error) {
	var modelPosts []model.Post
	err := r.DB.NewSelect().Model(&modelPosts).Order("created_at DESC").Limit(count).Scan(ctx)
	if err != nil {
		return nil, err
	}

	var posts entity.PostSlice
	for _, p := range modelPosts {
		posts = append(posts, p.ToPostEntity())
	}
	return posts, nil
}

func (r *PostRepository) Create(ctx context.Context, post *entity.Post) (*entity.Post, error) {
	modelPost := &model.Post{
		ID:        post.ID,
		Body:      post.Body,
		CreatedAt: post.CreatedAt,
	}

	_, err := r.DB.NewInsert().Model(modelPost).Exec(ctx)
	if err != nil {
		return nil, err
	}
	return modelPost.ToPostEntity(), nil
}

func (r *PostRepository) Delete(ctx context.Context, id string) error {
	post := &entity.Post{ID: id}
	_, err := r.DB.NewDelete().Model(post).Where("id = ?", id).Exec(ctx)
	return err
}
