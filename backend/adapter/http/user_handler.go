package http

import (
	"net/http"

	"github.com/saitamau-maximum/maxitter/backend/domain/entity"
	"github.com/saitamau-maximum/maxitter/backend/domain/service"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
)

type userHandler struct {
	userService service.IUserService
	Logger      echo.Logger
}

func NewUserHandler(service service.IUserService, logger echo.Logger) *userHandler {
	return &userHandler{
		userService: service,
		Logger:      logger,
	}
}

func (h *userHandler) FindAllUsers() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		user, err := h.userService.GetUsers(ctx)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusOK, user)
	}
}

func (h *userHandler) CreateUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		user := new(entity.User)
		if err := c.Bind(user); err != nil {
			h.Logger.Error(err)
			return c.JSON(500, err)
		}
		user, err := h.userService.CreateUser(ctx, user)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusOK, user)
	}
}

// func (sh *userHandler) FindUserByID() echo.HandlerFunc {
// 	return func(c echo.Context) error {
// 		ctx := c.Request().Context()
// 		userID, err := strconv.Atoi(c.Param("id"))
// 		if err != nil {
// 			return c.JSON(http.StatusBadRequest, err.Error())
// 		}
// 		user, err := sh.usecase.FindUserByID(ctx, userID)
// 		if err != nil {
// 			return c.JSON(http.StatusInternalServerError, err.Error())
// 		}
// 		return c.JSON(http.StatusOK, user)
// 	}
// }
