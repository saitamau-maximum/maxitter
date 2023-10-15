package main

import (
	"fmt"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	DB     *sqlx.DB
	Logger echo.Logger
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
	api.GET("/users", h.GetUsers)
	api.POST("/users/new", h.CreateUser)
	e.Logger.Fatal(e.Start(":8000"))
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

func (h *Handler) GetUsers(c echo.Context) error {
	users := []User{}
	err := h.DB.Select(&users, "SELECT * FROM users")
	if err != nil {
		h.Logger.Error(err)
		return c.JSON(500, err)
	}
	return c.JSON(200, users)
}

func (h *Handler) CreateUser(c echo.Context) error {
	id, err := uuid.NewRandom()
	if err != nil {
		h.Logger.Error(err)
		return c.JSON(500, err)
	}
	user := new(User)
	if err := c.Bind(user); err != nil {
		h.Logger.Error(err)
		return c.JSON(500, err)
	}
	user.ID = id.String()
	user.CreatedAt = time.Now().Format("2006-01-02 15:04:05")
	user.UpdatedAt = time.Now().Format("2006-01-02 15:04:05")

	_, err = h.DB.Exec("INSERT INTO users (id, username, created_at, updated_at, profile_image_url, bio) VALUES (?, ?, ?, ?, ?, ?)", user.ID, user.Name, user.CreatedAt, user.UpdatedAt, user.ProfileImageURL, user.Bio)
	if err != nil {
		h.Logger.Error(err)
		return c.JSON(500, err)
	}
	return c.JSON(200, user)
}
