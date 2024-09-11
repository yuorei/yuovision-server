package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/yuorei/video-server/db/sqlc"
)

type DB struct {
	Database *sqlc.Queries
}

func NewMySQLDB() *DB {
	MYSQL_HOST := os.Getenv("MYSQL_HOST")
	MYSQL_PORT := os.Getenv("MYSQL_PORT")
	MYSQL_DATABASE := os.Getenv("MYSQL_DATABASE")
	MYSQL_PASSWORD := os.Getenv("MYSQL_PASSWORD")
	MYSQL_USER := os.Getenv("MYSQL_USER")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", MYSQL_USER, MYSQL_PASSWORD, MYSQL_HOST, MYSQL_PORT, MYSQL_DATABASE)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Error: ", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal("Error: ", err)
	}

	queries := sqlc.New(db)

	return &DB{
		Database: queries,
	}
}
