package model

import (
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type Post struct {
	Id        int64     `bun:", pk, autoincrement"`
	Body      string    `bun:", notnull, varchar()"`
	CreatedAt time.Time `bun:", notnull, default:current_timestamp"`
}
