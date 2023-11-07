package infra

import (
	"database/sql"
	"fmt"

	"github.com/go-sql-driver/mysql"
	"github.com/saitamau-maximum/maxitter/backend/config"
)

func ConnectDB(cfg *config.Config) (*sql.DB, error) {
	cfg_mysql := cfg.MYSQL

	user := cfg_mysql.MYSQL_USER
	password := cfg_mysql.MYSQL_PASSWORD
	host := cfg_mysql.MYSQL_HOST
	port := cfg_mysql.MYSQL_PORT
	dbname := cfg_mysql.MYSQL_DATABASE

	c := mysql.Config{
		User:                 user,
		Passwd:               password,
		Net:                  "tcp",
		Addr:                 fmt.Sprintf("%s:%s", host, port),
		DBName:               dbname,
		AllowNativePasswords: true,
		ParseTime:            true,
		Collation:            "utf8mb4_unicode_ci",
	}

	db, err := sql.Open("mysql", c.FormatDSN())
	if err != nil {
		return nil, err
	}

	return db, nil
}
