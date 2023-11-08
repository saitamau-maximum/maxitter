package repository

import (
	"context"

	"github.com/saitamau-maximum/maxitter/backend/domain/entity"
	"github.com/saitamau-maximum/maxitter/backend/domain/repository"
	"github.com/saitamau-maximum/maxitter/backend/usecase/model"
	"github.com/uptrace/bun"
)

type UserRepository struct {
	DB *bun.DB
}

func NewUserRepository(db *bun.DB) repository.IUserService {
	return &UserRepository{DB: db}
}

func (r *UserRepository) GetAll(ctx context.Context) (entity.UserSlice, error) {
	var modelUsers []model.User
	err := r.DB.NewSelect().Model(&modelUsers).Scan(ctx)
	if err != nil {
		return nil, err
	}

	var users entity.UserSlice
	for _, u := range modelUsers {
		users = append(users, u.ToUserEntity())
	}

	return users, nil
}

func (r *UserRepository) GetByID(ctx context.Context, id string) (*entity.User, error) {
	user := &model.User{}
	err := r.DB.NewSelect().Model(user).Where("id = ?", id).Scan(ctx)
	if err != nil {
		return nil, err
	}
	return user.ToUserEntity(), nil
}
func (r *UserRepository) Create(ctx context.Context, user *entity.User) (*entity.User, error) {
	modelUser := &model.User{
		ID:           user.ID,
		Name:         user.Name,
		ProfileImage: user.ProfileImageURL,
		Bio:          user.Bio,
		CreatedAt:    user.CreatedAt,
		UpdatedAt:    user.UpdatedAt,
	}

	_, err := r.DB.NewInsert().Model(modelUser).Exec(ctx)
	if err != nil {
		return nil, err
	}

	return modelUser.ToUserEntity(), nil
}

func (r *UserRepository) Delete(ctx context.Context, id string) error {
	user := &model.User{ID: id}
	_, err := r.DB.NewDelete().Model(user).Where("id = ?", id).Exec(ctx)
	return err
}
