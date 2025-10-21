package database

import (
	"log/slog"

	"github.com/AC-Pcong/goscaff/pkg/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// NewDB creates a new GORM database instance.
// It is intended to be used as a wire provider.
func NewDB(cfg *config.Config, logger *slog.Logger) (*gorm.DB, func(), error) {
	dsn := cfg.Database.DSN

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, nil, err
	}

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(10)
	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(100)
	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	// sqlDB.SetConnMaxLifetime(time.Hour)

	logger.Info("Database connection established")

	cleanup := func() {
		logger.Info("Closing database connection...")
		if err := sqlDB.Close(); err != nil {
			logger.Error("Failed to close database connection", "error", err)
		}
	}

	return db, cleanup, nil
}
