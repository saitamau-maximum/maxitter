package model

import (
	"time"

	"github.com/saitamau-maximum/maxitter/backend/internal/entity"
	"github.com/uptrace/bun"
)

type Post struct {
	bun.BaseModel `bun:"posts,alias:posts"`
	ID            string    `bun:"id,type:varchar(36),nullzero,notnull"`
	Body          string    `bun:"body,type:text,nullzero,notnull"`
	CreatedAt     time.Time `bun:"created_at,nullzero,notnull"`
}

func (p *Post) ToPostEntity() *entity.Post {
	return entity.NewPostEntity(p.ID, p.Body, p.CreatedAt)
}
