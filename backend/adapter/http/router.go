package http

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	"github.com/uptrace/bun"

	"github.com/saitamau-maximum/maxitter/backend/config"
	"github.com/saitamau-maximum/maxitter/backend/domain/service"
	"github.com/saitamau-maximum/maxitter/backend/infra/repository"
)

const (
	apiRoot = "/api"

	// health check
	healthCheckRoot = "/health"

	// posts
	postsAPIRoot   = "/posts"
	postsPageParam = "page"

	// users
	usersAPIRoot          = "/users"
	createUserAPIRelative = "/new"
)

type Container struct {
	DB          *bun.DB
	UserService *service.UserService
	PostService *service.PostService
}

type Handler struct {
	Container Container
	Logger    echo.Logger
}

func InitRouter(cfg *config.Config, bunDB *bun.DB) *echo.Echo {
	e := echo.New()
	e.Debug = true
	e.Logger.SetLevel(0)
	// e.Use(
	// 	middleware.Logger(),
	// 	middleware.Recover(),
	// )

	userRepository := repository.NewUserRepository(bunDB)
	userService := service.NewUserService(userRepository)

	postRepository := repository.NewPostRepository(bunDB)
	postService := service.NewPostService(postRepository)

	container := Container{
		DB:          bunDB,
		UserService: userService,
		PostService: postService,
	}

	h := &Handler{Container: container, Logger: e.Logger}

	apiGroup := e.Group(apiRoot)
	// health check
	healthCheckGroup := apiGroup.Group(healthCheckRoot)
	{
		relativePath := ""
		healthCheckGroup.GET(relativePath, healthCheck)
	}

	userGroup := apiGroup.Group(usersAPIRoot)
	{
		handler := NewUserHandler(h.Container.UserService, h.Logger)
		// /users
		relativePath := ""
		userGroup.GET(relativePath, handler.FindAllUsers())
		// /users/new
		relativePath = createUserAPIRelative
		userGroup.POST(relativePath, handler.CreateUser())
	}

	postGroup := apiGroup.Group(postsAPIRoot)
	{
		handler := NewPostHandler(h.Container.PostService, h.Logger)
		// /posts
		relativePath := ""
		postGroup.GET(relativePath, handler.GetPosts())
		postGroup.POST(relativePath, handler.CreatePost())
	}

	return e
}
