package model

import (
	"time"

	"github.com/saitamau-maximum/maxitter/backend/internal/entity"
)

type User struct {
	ID              uint32    `bun:"id,pk,autoincrement"`
	Name            string    `bun:"username,unique,notnull"`
	ProfileImageURL string    `bun:"profile_image_url,notnull"`
	Bio             string    `bun:"bio,notnull,default:current_timestamp"`
	CreatedAt       time.Time `bun:"created_at,default:null"`
	UpdatedAt       time.Time `bun:"updated_at,notnull,default:current_timestamp"`
}

func (u *User) ToUserEntity() *entity.User {
	return entity.NewUserEntity(u.ID, u.Name, u.ProfileImageURL, u.Bio, u.CreatedAt, u.UpdatedAt)
}
