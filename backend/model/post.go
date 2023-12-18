package model

import (
	"time"

	"github.com/uptrace/bun"
)

type Post struct {
	bun.BaseModel `bun:"posts,alias:posts"`
	ID            string    `bun:"id,type:varchar(36),nullzero,notnull" json:"id"`
	Body          string    `bun:"body,type:text,nullzero,notnull" json:"body"`
	CreatedAt     time.Time `bun:"created_at,nullzero,notnull" json:"created_at"`
}
