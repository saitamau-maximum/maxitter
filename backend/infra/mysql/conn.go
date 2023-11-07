package infra

import (
	"database/sql"
	"fmt"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/mysqldialect"

	"github.com/saitamau-maximum/maxitter/backend/config"
)

func ConnectDB() *bun.DB {
	cfg := config.NewConfig()
	cfg_mysql := cfg.MYSQL

	user := cfg_mysql.MYSQL_USER
	password := cfg_mysql.MYSQL_PASSWORD
	host := cfg_mysql.MYSQL_HOST
	port := cfg_mysql.MYSQL_PORT
	dbname := cfg_mysql.MYSQL_DATABASE

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user,
		password,
		host,
		port,
		dbname,
	)

	con, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}

	db := bun.NewDB(con, mysqldialect.New())
	return db
}
