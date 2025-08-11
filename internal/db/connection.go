package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/upinmcSE/goshop/internal/config"
	"github.com/upinmcSE/goshop/internal/db/sqlc"
)

var DB sqlc.Querier
var DBConn *sql.DB

func InitDB() error {
	connStr := config.NewConfig().DNS()

	// Mở kết nối MySQL
	db, err := sql.Open("mysql", connStr)
	if err != nil {
		return fmt.Errorf("error opening DB connection: %v", err)
	}

	// Cấu hình pool
	db.SetMaxOpenConns(50)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(30 * time.Minute)
	db.SetConnMaxIdleTime(5 * time.Minute)

	// Kiểm tra kết nối với timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := db.PingContext(ctx); err != nil {
		return fmt.Errorf("db ping error: %v", err)
	}

	// Lưu connection và khởi tạo sqlc
	DBConn = db
	DB = sqlc.New(db)

	log.Println("Connected to MySQL")

	return nil
}
