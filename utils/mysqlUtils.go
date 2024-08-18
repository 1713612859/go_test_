package utils

import (
	"database/sql"
	"fmt"
	"time"

	//注册驱动器 _下划线表示执行驱动中的init函数，不使用其他函数
	_ "github.com/go-sql-driver/mysql"
)

func ConnectMysql() (*sql.DB, error) {
	// Use environment variables for sensitive information
	dbUser := "root"
	dbPass := "123lyl"
	dbHost := "127.0.0.1"
	dbName := "bill"

	if dbUser == "" || dbPass == "" || dbHost == "" || dbName == "" {
		return nil, fmt.Errorf("environment variables not set correctly")
	}

	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?parseTime=true", dbUser, dbPass, dbHost, dbName)

	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		return nil, fmt.Errorf("failed to open database connection: %w", err)
	}

	// Set connection pool parameters
	db.SetConnMaxIdleTime(10 * time.Minute)
	db.SetConnMaxLifetime(30 * time.Minute)
	db.SetMaxOpenConns(10)

	// Ping the database to check connectivity
	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	return db, nil
}
