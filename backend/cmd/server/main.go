package main

import (
	"context"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/mysqldialect"

	adapterHTTP "github.com/saitamau-maximum/maxitter/backend/adapter/http"
	"github.com/saitamau-maximum/maxitter/backend/config"
	infra "github.com/saitamau-maximum/maxitter/backend/infra/mysql"
	repo "github.com/saitamau-maximum/maxitter/backend/infra/repository"
)

const (
	PORT       = "8000"
	IMAGES_DIR = "./public/images"
)

type Container struct {
	DB             *bun.DB
	PostRepository *repo.PostRepository
	UserRepository *repo.UserRepository
}

type Handler struct {
	Container *Container
	Logger    echo.Logger
}

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	cfg := config.NewConfig()

	mySQLConn, err := infra.ConnectDB(cfg)
	bunDB := bun.NewDB(mySQLConn, mysqldialect.New())
	defer bunDB.Close()

	router := adapterHTTP.InitRouter(cfg, bunDB)

	srv := &http.Server{
		Addr:           ":" + PORT,
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		log.Println("server is running! addr: ", ":"+PORT)

		if err != nil {
			log.Fatalf("Failed to connect to MySQL: %+v", err)
		}

		if err := srv.ListenAndServe(); err != nil {
			log.Fatalf("Failed to listen and serve: %+v", err)
		}
	}()

	// Listen for the interrupt signal.
	<-ctx.Done()

	// Restore default behavior on the interrupt signal and notify user of shutdown.
	stop()
	log.Println("shutting down gracefully, press Ctrl+C again to force")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown: ", err)
	}

	log.Println("Server exiting")
}
