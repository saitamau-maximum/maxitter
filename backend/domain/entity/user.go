package entity

import (
	"time"
)

type User struct {
	ID              string    `db:"id" json:"id"`
	Name            string    `db:"username" json:"name"`
	ProfileImageURL string    `db:"profile_image_url" json:"profile_image_url"`
	Bio             string    `db:"bio" json:"bio"`
	CreatedAt       time.Time `db:"created_at" json:"created_at"`
	UpdatedAt       time.Time `db:"updated_at" json:"updated_at"`
}

type UserSlice []*User

func NewUserEntity(id string, name, profileImageURL, bio string, createdAt, updatedAt time.Time) *User {
	return &User{
		ID:              id,
		Name:            name,
		ProfileImageURL: profileImageURL,
		Bio:             bio,
		CreatedAt:       createdAt,
		UpdatedAt:       updatedAt,
	}
}
