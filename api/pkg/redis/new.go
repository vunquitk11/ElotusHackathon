package redis

import (
	"context"
	"fmt"
	"strconv"
	"time"

	redigo "github.com/gomodule/redigo/redis"
	"github.com/letsvote/api/pkg/app"
)

// NewClient connects to Redis and returns a client which is used to access Redis
// The projectName will be used from the app.Config to namespace the redis kv pairs by prefixing the key with `projectName:`
func NewClient(ctx context.Context, appCfg app.Config, url string, database int, opts ...Option) (*Client, error) {
	if err := appCfg.IsValid(); err != nil {
		return nil, err
	}

	pool := &redigo.Pool{
		MaxActive:       10,
		MaxIdle:         3,
		MaxConnLifetime: 9 * time.Minute, // https://docs.microsoft.com/en-us/azure/azure-cache-for-redis/cache-best-practices-connection#idle-timeout
		DialContext: func(ctx context.Context) (redigo.Conn, error) {
			return redigo.DialURLContext(
				ctx,
				url,
				redigo.DialDatabase(database),
			)
		},
	}

	cfg := config{
		pool: pool,
	}
	for _, opt := range opts {
		opt(&cfg)
	}

	client := Client{
		database:    strconv.Itoa(database),
		projectName: appCfg.Project,
		pool:        pool,
	}

	if cfg.pingUponInit {
		if err := client.Ping(ctx); err != nil {
			return nil, err
		}
		fmt.Println("Redis ping successful")
	}

	fmt.Println("Redis initialized")

	return &client, nil
}
