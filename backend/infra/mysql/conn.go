package infra

import (
	"database/sql"
	"fmt"

	"github.com/saitamau-maximum/maxitter/backend/config"
)

func ConnectDB(cfg *config.Config) (*sql.DB, error) {
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

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	return db, nil
}
