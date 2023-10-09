package main

import (
	"fmt"
	"log"
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
