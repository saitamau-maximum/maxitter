package entity

import (
	"time"
)

type Post struct {
	ID        string    `db:"id" json:"id"`
	Body      string    `db:"body" json:"body"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
}

type PostSlice []*Post

func NewPostEntity(id string, body string, createdAt time.Time) *Post {
	return &Post{
		ID:        id,
		Body:      body,
		CreatedAt: createdAt,
	}
}
