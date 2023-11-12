package redis

import (
	"errors"

	redigo "github.com/gomodule/redigo/redis"
)

var (
	// ErrNilReply is the constant to indicate redis nil reply
	// ErrNil indicates that a reply value is nil
	ErrNilReply = redigo.ErrNil

	// ErrSetFailed means the set command did not return an OK response
	ErrSetFailed = errors.New("set command failed")
)
