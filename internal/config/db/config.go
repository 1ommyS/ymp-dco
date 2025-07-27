package db

import (
	"context"
	"fmt"
	"github.com/itpark/market/dco/internal/config"
	"github.com/itpark/market/dco/internal/telemetry/logging"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type DbConnection struct {
	DB *sqlx.DB
}

func CreateConnectionString(cfg *config.Config) string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.DB.Host,
		cfg.DB.Port,
		cfg.DB.Login,
		cfg.DB.Password,
		cfg.DB.DbName,
	)
}

func InitConnection(ctx context.Context, cfg *config.Config) *DbConnection {
	db, err := sqlx.ConnectContext(ctx, "postgres", CreateConnectionString(cfg))
	if err != nil {
		logging.Error(ctx, " Error creating connection pool: %v", err.Error())
	}
	if cfg.DB.MaxPoolConnections > 0 {
		db.SetMaxOpenConns(cfg.DB.MaxPoolConnections)
	}
	return &DbConnection{
		DB: db,
	}
}
