package model

import (
	"time"

	"github.com/uptrace/bun"
)

type User struct {
	bun.BaseModel `bun:"users,alias:users"`
	ID            string    `bun:"id,type:varchar(36),nullzero,notnull" json:"id"`
	Name          string    `bun:"name,type:varchar(255),nullzero,notnull" json:"name"`
	ProfileImage  string    `bun:"profile_image_url,type:varchar(255),nullzero" json:"profile_image_url"`
	Bio           string    `bun:"bio,type:varchar(1024),nullzero" json:"bio"`
	CreatedAt     time.Time `bun:"created_at,nullzero,notnull" json:"created_at"`
	UpdatedAt     time.Time `bun:"updated_at,nullzero,notnull" json:"updated_at"`
}
