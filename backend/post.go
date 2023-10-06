package main

import (
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"

	"github.com/saitamau-maximum/maxitter/backend/webhook"
)

const (
	USERNAME   = "Maxitter"
	AVATER_URL = ""
)

type Post struct {
	ID        string `db:"id" json:"id"`
	Body      string `db:"body" json:"body"`
	CreatedAt string `db:"created_at" json:"created_at"`
}

func sendPostWebhook(post *Post) bool {
	whurl := getEnv("DISCORD_WEBHOOK_URL", "")

	if whurl == "" {
		log.Println("DISCORD_WEBHOOK_URL is empty")
		return false
	}

	dw := &webhook.DiscordWebhook{
		UserName:  USERNAME,
		AvatarURL: AVATER_URL,
		Content:   "",
		Embeds: []webhook.DiscordEmbed{
			{
				Title: "",
				Desc:  post.Body,
				URL:   "https://maxitter2.netlify.app/",
				Color: 0x23d9eb,
				Author: webhook.DiscordAuthor{
					Name: USERNAME,
					Icon: AVATER_URL,
				},
				TimeStamp: post.CreatedAt,
			},
		},
	}

	result := webhook.SendWebhook(whurl, USERNAME, AVATER_URL, dw)

	if !result {
		log.Println("SendWebhook error")
		return false
	}

	return true
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

	if !sendPostWebhook(post) {
		log.Println("sendPostWebhook error")
	}

	return c.JSON(200, post)
}
