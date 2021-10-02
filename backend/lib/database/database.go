package database

import (
	"context"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

// Database is a connection to a PostgreSQL database
type Database struct {
	pool *pgxpool.Pool
}

// New creates a new Database. It reads all configuration from PG* environment variables
func New() (*Database, error) {
	// Read configuration from environment variables
	config, err := pgxpool.ParseConfig("")
	if err != nil {
		return nil, err
	}
	config.LazyConnect = true

	db, err := pgxpool.ConnectConfig(context.Background(), config)
	if err != nil {
		return nil, err
	}

	return &Database{pool: db}, nil
}

// Healthy returns an error if performing a simple query to database fails
func (db *Database) Healthy(ctx context.Context) error {
	conn, err := db.pool.Acquire(ctx)
	if err != nil {
		return err
	}
	defer conn.Release()

	row := conn.QueryRow(ctx, "SELECT 1")

	var result int
	if err := row.Scan(&result); err != nil {
		return err
	}
	return nil
}

// Conn returns a new connection from database connection pool
func (db *Database) Conn(ctx context.Context) (*pgxpool.Conn, error) {
	return db.pool.Acquire(ctx)
}

// Exec runs a single query that doesn't return a result. Mainly to be used for UPDATEs and DELETES
func (db *Database) Exec(ctx context.Context, query string, v ...interface{}) (pgconn.CommandTag, error) {
	return db.pool.Exec(ctx, query, v...)
}

// Tx runs given function inside a transaction. Transaction is rolled back if the function returns an error
func (db *Database) Tx(ctx context.Context, f func(ctx context.Context, tx pgx.Tx) error) error {
	tx, err := db.pool.Begin(ctx)
	if err != nil {
		return err
	}
	defer func() {
		_ = tx.Rollback(ctx)
	}()

	if err := f(ctx, tx); err != nil {
		return err
	}
	return tx.Commit(ctx)
}
