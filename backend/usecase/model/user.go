package model

import (
	"time"

	"github.com/uptrace/bun"

	"github.com/saitamau-maximum/maxitter/backend/domain/entity"
)

type User struct {
	bun.BaseModel `bun:"users,alias:users"`
	ID            string    `bun:"id,type:varchar(36),nullzero,notnull"`
	Name          string    `bun:"username,type:varchar(255),nullzero,notnull"`
	ProfileImage  string    `bun:"profile_image_url,type:varchar(255),nullzero"`
	Bio           string    `bun:"bio,type:varchar(1024),nullzero"`
	CreatedAt     time.Time `bun:"created_at,nullzero,notnull"`
	UpdatedAt     time.Time `bun:"updated_at,nullzero,notnull"`
}

func (u *User) ToUserEntity() *entity.User {
	return entity.NewUserEntity(u.ID, u.Name, u.ProfileImage, u.Bio, u.CreatedAt, u.UpdatedAt)
}
