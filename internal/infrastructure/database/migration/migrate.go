package main

import (
	"context"
	"dropx/internal/domain/models"
	"dropx/internal/infrastructure/database"
	"dropx/pkg/config"
	"dropx/pkg/logger"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"gorm.io/gorm"
	"os"
	"strconv"
	"strings"
)

const (
	SQL_SOURCE_URL       = "file://./internal/infrastructure/database/migration/migrations"
	FAILED_MIGRATION_LOG = "Failed to migrate"
)

func main() {
	ctx := context.Background()
	db := database.PostgresqlDB(ctx)
	logger.InfoLog("Start migration", nil)
	migratePostgres(ctx, db)
	//MigrateSqlScript(db)
	logger.InfoLog("End migration", nil)
}

func migratePostgres(ctx context.Context, db *gorm.DB) {
	err := AddExtension(ctx, db)
	if err != nil {
		logger.FatalLog(fmt.Errorf("-------- Migration Postgres error: %s ", err.Error()), nil)
	}

	err = db.Migrator().AutoMigrate(
		&models.User{},
		&models.Warehouse{},
		&models.Product{},
		&models.ProductBundle{},
		&models.WarehouseOrder{},
		&models.WarehouseOrderItem{},
		&models.WarehouseInventory{},
	)
	if err != nil {
		logger.FatalLog(fmt.Errorf("Failed to migrate: %v ", err), nil)
	}

	logger.InfoLog("End migration", nil)
}

func MigrateSqlScript(db *gorm.DB) {
	m, err := migrate.New(
		SQL_SOURCE_URL,
		getPostUrl())
	if err != nil {
		logger.FatalLog(fmt.Errorf("%s: %v", FAILED_MIGRATION_LOG, err), nil)
	}
	args := os.Args[1:]
	if len(args) > 0 && args[0] == "down" {
		if len(args) > 1 {
			if myStep, err := strconv.Atoi(args[1]); err == nil {
				m.Steps(myStep)
			}
		}
		if err = m.Down(); err != nil {
			logger.ErrorLog(fmt.Errorf("%s: %v", FAILED_MIGRATION_LOG, err), nil)
		}
	} else {
		if err = m.Up(); err != nil {
			logger.ErrorLog(fmt.Errorf("%s: %v", FAILED_MIGRATION_LOG, err), nil)
		}
	}
	logger.InfoLog("End migration", nil)
}

func getPostUrl() string {
	var url string
	splitedStr := strings.Split(config.Global.PostgresURL, " ")
	if len(splitedStr) >= 5 {
		url = fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable&TimeZone=Asia/Tokyo",
			strings.Replace(splitedStr[1], "user=", "", 1),
			strings.Replace(splitedStr[2], "password=", "", 1),
			strings.Replace(splitedStr[0], "host=", "", 1),
			strings.Replace(splitedStr[4], "port=", "", 1),
			strings.Replace(splitedStr[3], "dbname=", "", 1),
		)
	}
	return url
}
func AddExtension(ctx context.Context, db *gorm.DB) error {
	err := db.Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp"`).Error
	return err
}
