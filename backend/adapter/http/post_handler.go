package http

import (
	"net/http"
	"strconv"

	"github.com/saitamau-maximum/maxitter/backend/domain/entity"
	"github.com/saitamau-maximum/maxitter/backend/domain/service"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
)

type postHandler struct {
	postService service.IPostService
	Logger      echo.Logger
}

func NewPostHandler(service service.IPostService, logger echo.Logger) *postHandler {
	return &postHandler{
		postService: service,
		Logger:      logger,
	}
}

func (h *postHandler) GetPosts() echo.HandlerFunc {
	return func(c echo.Context) error {
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
		posts, err := h.postService.GetRecentPosts(ctx, int(index))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusOK, posts)
	}
}

func (h *postHandler) CreatePost() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		post := new(entity.Post)
		if err := c.Bind(post); err != nil {
			h.Logger.Error(err)
			return c.JSON(500, err)
		}
		post, err := h.postService.CreatePost(ctx, post)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusOK, post)
	}
}

// func (sh *postHandler) FindPostByID() echo.HandlerFunc {
// 	return func(c echo.Context) error {
// 		ctx := c.Request().Context()
// 		postID, err := strconv.Atoi(c.Param("id"))
// 		if err != nil {
// 			return c.JSON(http.StatusBadRequest, err.Error())
// 		}
// 		post, err := sh.usecase.FindPostByID(ctx, postID)
// 		if err != nil {
// 			return c.JSON(http.StatusInternalServerError, err.Error())
// 		}
// 		return c.JSON(http.StatusOK, post)
// 	}
// }
