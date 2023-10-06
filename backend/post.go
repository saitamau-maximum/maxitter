package main

import (
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type Post struct {
	ID        string `db:"id" json:"id"`
	Body      string `db:"body" json:"body"`
	CreatedAt string `db:"created_at" json:"created_at"`
}

func (h *Handler) GetPosts(c echo.Context) error {
	posts := []Post{}
	err := h.DB.Select(&posts, "SELECT * FROM posts ORDER BY created_at DESC LIMIT 20")
	if err != nil {
		h.Logger.Error(err)
		return c.JSON(500, err)
	}
	return c.JSON(200, posts)
}

func (h *Handler) CreatePost(c echo.Context) error {
	id, err := uuid.NewRandom()
	if err != nil {
		h.Logger.Error(err)
		return c.JSON(500, err)
	}
	post := new(Post)
	if err := c.Bind(post); err != nil {
		h.Logger.Error(err)
		return c.JSON(500, err)
	}
	post.ID = id.String()
	post.CreatedAt = time.Now().Format("2006-01-02 15:04:05")

	_, err = h.DB.Exec("INSERT INTO posts (id, body, created_at) VALUES (?, ?, ?)", post.ID, post.Body, post.CreatedAt)
	if err != nil {
		h.Logger.Error(err)
		return c.JSON(500, err)
	}
	return c.JSON(200, post)
}
