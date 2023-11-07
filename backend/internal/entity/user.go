package entity

import (
	"time"
)

type User struct {
	ID              string    `bun:"id" json:"id"`
	Name            string    `bun:"username" json:"name"`
	ProfileImageURL string    `bun:"profile_image_url" json:"profile_image_url"`
	Bio             string    `bun:"bio" json:"bio"`
	CreatedAt       time.Time `bun:"created_at" json:"created_at"`
	UpdatedAt       time.Time `bun:"updated_at" json:"updated_at"`
}

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
