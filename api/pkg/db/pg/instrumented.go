package pg

import (
	"context"
	"database/sql"
)

// instrumentedDB wraps the *sql.DB provided
type instrumentedDB struct {
	*sql.DB
	info InstanceInfo
}

// BeginTx begins a transaction with the database in receiver and returns a Transactor
func (i instrumentedDB) BeginTx(ctx context.Context, opts *sql.TxOptions) (Transactor, error) {
	return i.DB.BeginTx(ctx, opts)
}

// ExecContext wraps the base connector
func (i instrumentedDB) ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	return i.DB.ExecContext(ctx, query, args...)
}

// QueryContext wraps the base connector
func (i instrumentedDB) QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error) {
	return i.DB.QueryContext(ctx, query, args...)
}

// QueryRowContext wraps the base connector
func (i instrumentedDB) QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row {
	return i.DB.QueryRowContext(ctx, query, args...)
}

// InstanceInfo returns info about the DB
func (i instrumentedDB) InstanceInfo() InstanceInfo {
	return i.info
}

// instrumentedTx wraps the Transactor provided
type instrumentedTx struct {
	Transactor
	info InstanceInfo
	ctx  context.Context
}

// ExecContext wraps the base connector
func (i instrumentedTx) ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	return i.Transactor.ExecContext(ctx, query, args...)
}

// QueryContext wraps the base connector
func (i instrumentedTx) QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error) {
	return i.Transactor.QueryContext(ctx, query, args...)
}

// QueryRowContext wraps the base connector
func (i instrumentedTx) QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row {
	return i.Transactor.QueryRowContext(ctx, query, args...)
}

// Commit commits the transaction
func (i instrumentedTx) Commit() error {
	return i.Transactor.Commit()
}

// Rollback aborts the transaction
func (i instrumentedTx) Rollback() error {
	return i.Transactor.Rollback()
}
