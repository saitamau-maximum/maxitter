package handler

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/uptrace/bun"

	"github.com/saitamau-maximum/maxitter/backend/model"
)

type Handler struct {
	DB     *bun.DB
	Logger echo.Logger
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

	ctx := c.Request().Context()

	modelPosts := []model.Post{}

	err = h.DB.NewSelect().Model(&modelPosts).Order("created_at DESC").Limit(20).Offset(int(index)).Scan(ctx)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	fmt.Printf("%#v\n", modelPosts)

	return c.JSON(http.StatusOK, modelPosts)
}

func (h *Handler) CreatePost(c echo.Context) error {
	id, err := uuid.NewRandom()
	if err != nil {
		h.Logger.Error(err)
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	post := new(model.Post)
	if err := c.Bind(post); err != nil {
		h.Logger.Error(err)
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	post.ID = id.String()
	post.CreatedAt = time.Now().Round(time.Millisecond)

	modelPost := &model.Post{
		ID:        post.ID,
		Body:      post.Body,
		CreatedAt: post.CreatedAt,
	}

	ctx := c.Request().Context()

	_, err = h.DB.NewInsert().Model(modelPost).Exec(ctx)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, modelPost)
}

func (h *Handler) GetUsers(c echo.Context) error {
	ctx := c.Request().Context()

	modelUsers := []model.User{}
	err := h.DB.NewSelect().Model(&modelUsers).Scan(ctx)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, modelUsers)
}

func (h *Handler) CreateUser(c echo.Context) error {
	id, err := uuid.NewRandom()
	if err != nil {
		h.Logger.Error(err)
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	user := new(model.User)
	if err := c.Bind(user); err != nil {
		h.Logger.Error(err)
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	user.ID = id.String()

	user.CreatedAt = time.Now().Round(time.Millisecond)
	user.UpdatedAt = time.Now().Round(time.Millisecond)

	modelUser := &model.User{
		ID:           user.ID,
		Name:         user.Name,
		ProfileImage: user.ProfileImage,
		Bio:          user.Bio,
		CreatedAt:    user.CreatedAt,
		UpdatedAt:    user.UpdatedAt,
	}

	ctx := c.Request().Context()

	_, err = h.DB.NewInsert().Model(modelUser).Exec(ctx)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, modelUser)
}
