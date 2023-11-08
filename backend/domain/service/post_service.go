package service

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/saitamau-maximum/maxitter/backend/domain/entity"
	"github.com/saitamau-maximum/maxitter/backend/domain/repository"
)

type IPostService interface {
	GetPosts(ctx context.Context) (entity.PostSlice, error)
	GetRecentPosts(ctx context.Context, count int) (entity.PostSlice, error)
	CreatePost(ctx context.Context, post *entity.Post) (*entity.Post, error)
}

type PostService struct {
	PostRepository repository.IPostRepository
}

func NewPostService(repo repository.IPostRepository) *PostService {
	return &PostService{PostRepository: repo}
}

func (s *PostService) GetPosts(ctx context.Context) (entity.PostSlice, error) {
	posts, err := s.PostRepository.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	return posts, nil
}

func (s *PostService) GetRecentPosts(ctx context.Context, count int) (entity.PostSlice, error) {
	posts, err := s.PostRepository.GetRecent(ctx, count)
	if err != nil {
		return nil, err
	}
	return posts, nil
}

func (s *PostService) CreatePost(ctx context.Context, post *entity.Post) (*entity.Post, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	post.ID = id.String()
	post.CreatedAt = time.Now()

	_, err = s.PostRepository.Create(ctx, post)
	if err != nil {
		return nil, err
	}

	return post, nil
}
