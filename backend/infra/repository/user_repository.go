package repository

import (
	"context"

	"github.com/uptrace/bun"

	"github.com/saitamau-maximum/maxitter/backend/internal/entity"
)

type UserRepository struct {
	db *bun.DB
}

func NewUserRepository(db *bun.DB) *UserRepository {
	return &UserRepository{
		db,
	}
}

func (r *UserRepository) Create(ctx context.Context, e *entity.User) (string, error) {
	_, err := r.db.NewInsert().Model(e).Exec(ctx)
	if err != nil {
		return e.ID, err
	}
	return e.ID, nil
}

func (r *UserRepository) Find(ctx context.Context, id string) (*entity.User, error) {
	user := &entity.User{}
	err := r.db.NewSelect().Model(user).Where("id = ?", id).Scan(ctx)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepository) List(ctx context.Context) ([]*entity.User, error) {
	var users []*entity.User
	err := r.db.NewSelect().Model(&users).Scan(ctx)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (r *UserRepository) Delete(ctx context.Context, id string) error {
	_, err := r.db.NewDelete().Model(&entity.User{}).Where("id = ?", id).Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}
