package main

import (
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	DB     *sqlx.DB
	Logger echo.Logger
}

var (
	SQL_PATH = "./sql"
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
	e.Logger.Fatal(e.Start(":8000"))
}
