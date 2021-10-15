package database

import (
	"context"
	"fmt"
	"strings"

	pgxpool "github.com/jackc/pgx/v4/pgxpool"
	"github.com/jlpadilla/search-indexer/pkg/config"
	"k8s.io/klog/v2"
)

var pool *pgxpool.Pool

func init() {
	klog.Info("Initializing database connection.")
	initializePool()
}

func initializePool() {
	cfg := config.New()

	// TODO: Validate configuration

	database_url := fmt.Sprintf("postgresql://%s:%s@%s:%d/%s", cfg.DBUser, cfg.DBPass, cfg.DBHost, cfg.DBPort, cfg.DBName)
	klog.Info("Connecting to PostgreSQL at: ", strings.ReplaceAll(database_url, cfg.DBPass, "*****"))
	config, connerr := pgxpool.ParseConfig(database_url)
	if connerr != nil {
		klog.Info("Error connecting to DB:", connerr)
	}
	// config.MaxConns = maxConnections
	conn, err := pgxpool.ConnectConfig(context.Background(), config)

	if err != nil {
		klog.Error("Unable to connect to database: %v\n", err)
	}

	pool = conn
}

func GetConnection() *pgxpool.Pool {
	err := pool.Ping(context.Background())
	if err != nil {
		panic(err)
	}
	klog.Info("Successfully connected to database!")
	return pool
}
