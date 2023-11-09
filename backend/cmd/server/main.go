package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	"github.com/saitamau-maximum/maxitter/backend/handler"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/mysqldialect"
)

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

func connectDB() (*sql.DB, error) {
	user := getEnv("MYSQL_USER", "user")
	password := getEnv("MYSQL_PASSWORD", "password")
	host := getEnv("MYSQL_HOST", "database")
	port := getEnv("MYSQL_PORT", "3306")
	dbname := getEnv("MYSQL_DATABASE", "db")

	c := mysql.Config{
		User:   user,
		Passwd: password,
		Net:    "tcp",
		Addr:   fmt.Sprintf("%s:%s", host, port),
		DBName: dbname,
	}

	db, err := sql.Open("mysql", c.FormatDSN())
	if err != nil {
		return nil, err
	}

	return db, nil
}

func main() {
	e := echo.New()
	e.Debug = true
	e.Logger.SetLevel(0)

	db, err := connectDB()
	if err != nil {
		e.Logger.Fatal(err)
	}

	bunDB := bun.NewDB(db, mysqldialect.New())
	defer bunDB.Close()

	h := &handler.Handler{DB: bunDB, Logger: e.Logger}
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
