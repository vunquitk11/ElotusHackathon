package pg

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/stdlib"
	"github.com/letsvote/api/pkg/app"
	pkgerrors "github.com/pkg/errors"
)

// NewPool opens a new DB connection pool, pings it and returns the pool
func NewPool(
	ctx context.Context,
	appCfg app.Config,
	url string,
	poolMaxOpenConns int,
	poolMaxIdleConns int,
	opts ...Option,
) (BeginnerExecutor, error) {

	if err := appCfg.IsValid(); err != nil {
		return nil, err
	}

	connCfg, err := pgx.ParseConfig(url)
	if err != nil {
		return nil, pkgerrors.WithStack(fmt.Errorf("parsing pgx config failed. err: %w", err))
	}
	connStr := stdlib.RegisterConnConfig(connCfg)

	pool, err := sql.Open("pgx", connStr)
	if err != nil {
		return nil, pkgerrors.WithStack(fmt.Errorf("opening DB failed. err: %w", err))
	}
	pool.SetConnMaxLifetime(29 * time.Minute) // Azure's default is 30 mins.
	pool.SetMaxOpenConns(poolMaxOpenConns)
	pool.SetMaxIdleConns(poolMaxIdleConns)
	cfg := config{
		pgxCfg: connCfg,
		pool:   pool,
	}
	for _, opt := range opts {
		opt(&cfg)
	}

	info := InstanceInfo{
		dbName: connCfg.Database,
	}

	if cfg.pingUponInit {
		if err = pool.PingContext(ctx); err != nil {
			return nil, pkgerrors.WithStack(fmt.Errorf("unable to ping DB. err: %w", err))
		}
	}

	return &instrumentedDB{
		DB:   pool,
		info: info,
	}, nil
}
