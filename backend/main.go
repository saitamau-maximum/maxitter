package main

import (
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"

	"github.com/saitamau-maximum/maxitter/backend/external"
)

type Handler struct {
	DB     *sqlx.DB
	Logger echo.Logger
}

type Post struct {
	ID        string `db:"id" json:"id"`
	Body      string `db:"body" json:"body"`
	CreatedAt string `db:"created_at" json:"created_at"`
}

var (
	SQL_PATH   = "./sql"
	IMAGES_DIR = "./public/images"
)

func getEnv(key, fallback string) string {
	value, ok := os.LookupEnv(key)
	if !ok {
		value = fallback
	}
	return value
}

func connectDB() *sqlx.DB {
	user := getEnv("MYSQL_USER", "user")
	password := getEnv("MYSQL_PASSWORD", "password")
	host := getEnv("MYSQL_HOST", "database")
	port := getEnv("MYSQL_PORT", "3306")
	dbname := getEnv("MYSQL_DATABASE", "db")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, port, dbname)

	con, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		panic(err)
	}
	return con
}

func migrate() {
	log.Println("migrate start")
	db := connectDB()
	defer db.Close()

	files, err := os.ReadDir(SQL_PATH)
	if err != nil {
		panic(err)
	}
	log.Println("migrate files: ", files)

	for _, file := range files {
		log.Println("migrate: " + file.Name())
		data, err := os.ReadFile(SQL_PATH + "/" + file.Name())
		if err != nil {
			panic(err)
		}
		_, err = db.Exec(string(data))
		if err != nil {
			panic(err)
		}
	}

	// indexをposts.created_atにつける
	_, err = db.Exec("CREATE INDEX posts_latest_idx ON posts (created_at)")
	if err != nil {
		log.Println("index already exists")
	}

	log.Println("migrate end")
}

func init() {
	migrate()
}

func main() {
	e := echo.New()
	e.Debug = true
	e.Logger.SetLevel(0)
	db := connectDB()
	defer db.Close()
	h := &Handler{DB: db, Logger: e.Logger}
	api := e.Group("/api")
	api.GET("/posts", h.GetPosts)
	api.POST("/posts", h.CreatePost)
	api.GET("/health", func(c echo.Context) error {
		e.Logger.Info("health check")
		return c.JSON(200, "ok")
	})
	e.Logger.Fatal(e.Start(":8000"))
}

const (
	DISCORD_USERNAME   = "Maxitter 投稿通知"
	DISCORD_AVATAR_URL = ""
)

func sendPostWebhookDiscord(post *Post) error {
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

	if err := sendPostWebhookDiscord(post); err != nil {
		// Discord Webhook送信に失敗してもエラーにしない
		fmt.Printf("sendPostWebhook error: %v", err)
	}

	return c.JSON(200, post)
}
