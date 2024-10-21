package Commands

import (
	"github.com/codecrafters-io/redis-starter-go/app/Storage"
	"github.com/codecrafters-io/redis-starter-go/app/api"
	"net"
	"time"
)

type OptionFunc func(*SetArgs)

func WithTTL(ttl time.Duration) OptionFunc {
	return func(s *SetArgs) {
		s.ttl = ttl
	}
}

type SetArgs struct {
	key   string
	value string
	GET   bool
	ttl   time.Duration
}

func SET(conn net.Conn, key string, value string, opts ...OptionFunc) {
	args := &SetArgs{
		key:   key,
		value: value,
	}

	for _, opt := range opts {
		opt(args)
	}

	storage := Storage.GetInstance()
	storage.Set(args.key, args.value)

	api.OK(conn)
}
