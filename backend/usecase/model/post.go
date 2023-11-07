package model

import (
	"time"

	"github.com/saitamau-maximum/maxitter/backend/internal/entity"
)

type Post struct {
	ID        uint32    `bun:"id,pk,autoincrement"`
	Body      string    `bun:"body,unique,notnull"`
	CreatedAt time.Time `bun:"created_at,default:null"`
	// UpdatedAt       time.Time `bun:"updated_at,notnull,default:current_timestamp"`
}

func (p *Post) ToPostEntity() *entity.Post {
	return entity.NewPostEntity(p.ID, p.Body, p.CreatedAt)
}
