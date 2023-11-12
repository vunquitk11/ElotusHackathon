package redis

import (
	"time"

	redigo "github.com/gomodule/redigo/redis"
)

type config struct {
	pool         *redigo.Pool
	pingUponInit bool
}

// Option is an optional config used to modify the client's behaviour
type Option func(*config)

// PoolMaxOpenConns sets the max open connections for the pool
func PoolMaxOpenConns(num int) Option {
	return func(c *config) {
		c.pool.MaxActive = num
	}
}

// PoolMaxIdleConns sets the max idle connections for the pool
func PoolMaxIdleConns(num int) Option {
	return func(c *config) {
		c.pool.MaxIdle = num
	}
}

// PoolMaxConnLifetime sets the max duration a connection should be kept alive in the pool
func PoolMaxConnLifetime(v time.Duration) Option {
	return func(c *config) {
		c.pool.MaxConnLifetime = v
	}
}

// AttemptPingUponStartup will ping the Redis instance upon startup. If this fails, NewClient will return error
func AttemptPingUponStartup() Option {
	return func(c *config) {
		c.pingUponInit = true
	}
}
