package main

import (
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"

	"github.com/saitamau-maximum/maxitter/backend/external"
)

const (
	DISCORD_USERNAME   = "Maxitter 投稿通知"
	DISCORD_AVATAR_URL = ""
)

type Post struct {
	ID        string `db:"id" json:"id"`
	Body      string `db:"body" json:"body"`
	CreatedAt string `db:"created_at" json:"created_at"`
}

func sendPostWebhook(post *Post) error {
	discord_webhook_url := getEnv("DISCORD_WEBHOOK_URL", "")

	if discord_webhook_url == "" {
		return fmt.Errorf("DISCORD_WEBHOOK_URL is empty")
	}

	discord_webhook := &external.DiscordWebhook{
		UserName:  DISCORD_USERNAME,
		AvatarURL: DISCORD_AVATAR_URL,
		Content:   "",
		Embeds: []external.DiscordEmbed{
			{
				Title: "",
				Desc:  post.Body,
				URL:   "",
				Color: 0x23d9eb,
				Author: external.DiscordAuthor{
					Name: "匿名のユーザー",
					Icon: DISCORD_AVATAR_URL,
				},
				TimeStamp: post.CreatedAt,
			},
		},
	}

	result := external.SendWebhookDiscord(
		discord_webhook_url,
		DISCORD_USERNAME,
		DISCORD_AVATAR_URL,
		discord_webhook,
	)

	if result != nil {
		return fmt.Errorf("sendWebhook error: %v", result)
	}

	return nil
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

	if err := sendPostWebhook(post); err != nil {
		return fmt.Errorf("sendPostWebhook error: %v", err)
	}

	return c.JSON(200, post)
}
