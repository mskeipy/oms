package database

import (
	"context"
	"dropx/pkg/config"
	"dropx/pkg/logger"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"sync"
	"time"
)

var db *gorm.DB
var once sync.Once

func init() {
	config.InitApp()
}

func PostgresqlDB(ctx context.Context) *gorm.DB {
	once.Do(func() {
		db = InitDB()
	})
	return db.WithContext(ctx)
}

func InitDB() *gorm.DB {
	logger.InfoLog("Connecting to postgresql!", nil)
	driver := postgres.Open(config.Global.PostgresURL)
	conn, err := gorm.Open(driver, &gorm.Config{
		Logger: logger.NewQueryLogger(500 * time.Millisecond),
	})
	if err != nil {
		logger.FatalLog(fmt.Errorf("Failed to open PostgreSQL database driver: "+err.Error()), nil)
		return nil
	}

	sqlDB, err := conn.DB()
	if err != nil {
		logger.FatalLog(fmt.Errorf("Failed to connect to the PostgreSQL database: "+err.Error()), nil)
		return nil
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	logger.InfoLog("Connected to PostgreSQL successfully!", nil)
	return conn
}

func Transaction(ctx context.Context, callback func(tx *gorm.DB) error) (err error) {
	logger.TransactionLogger(ctx, "BEGIN TRANSACTION")
	tx := PostgresqlDB(ctx).Begin()
	err = callback(tx)

	if err != nil {
		tx.Rollback()
		logger.TransactionLogger(ctx, "ROLLBACK")
		return err
	}

	tx.Commit()
	logger.TransactionLogger(ctx, "COMMIT")
	return nil
}
