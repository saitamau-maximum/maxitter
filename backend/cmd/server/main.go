package main

import (
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/uptrace/bun"

	infra "github.com/saitamau-maximum/maxitter/backend/infra/mysql"
	repo "github.com/saitamau-maximum/maxitter/backend/infra/repository"
	"github.com/saitamau-maximum/maxitter/backend/internal/entity"
)

const IMAGES_DIR = "./public/images"

type Container struct {
	DB             *bun.DB
	PostRepository *repo.PostRepository
	UserRepository *repo.UserRepository
}

type Handler struct {
	Container *Container
	Logger    echo.Logger
}

func init() {
	// migrate()
}

func main() {
	e := echo.New()
	e.Debug = true
	e.Logger.SetLevel(0)
	db := infra.ConnectDB()
	defer db.Close()

	container := Container{
		DB:             db,
		PostRepository: repo.NewPostRepository(db),
		UserRepository: repo.NewUserRepository(db),
	}

	h := &Handler{Container: &container, Logger: e.Logger}

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
	pageParam := c.QueryParam("page")
	if pageParam == "" {
		pageParam = "0"
	}
	page, err := strconv.ParseUint(pageParam, 10, 0)
	if err != nil {
		return c.JSON(400, err)
	}

	index := page * 20

	posts, err := h.Container.PostRepository.GetRecentPosts(c.Request().Context(), int(index))
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
	post := new(entity.Post)
	if err := c.Bind(post); err != nil {
		h.Logger.Error(err)
		return c.JSON(500, err)
	}
	post.ID = id.ID()
	post.CreatedAt = time.Now()

	_, err = h.Container.PostRepository.Create(post)
	if err != nil {
		h.Logger.Error(err)
		return c.JSON(500, err)
	}
	return c.JSON(200, post)
}

func (h *Handler) GetUsers(c echo.Context) error {
	users, err := h.Container.UserRepository.List(c.Request().Context())
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
	user := new(entity.User)
	if err := c.Bind(user); err != nil {
		h.Logger.Error(err)
		return c.JSON(500, err)
	}
	user.ID = id.ID()
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	_, err = h.Container.UserRepository.Create(user)
	if err != nil {
		h.Logger.Error(err)
		return c.JSON(500, err)
	}
	return c.JSON(200, user)
}
