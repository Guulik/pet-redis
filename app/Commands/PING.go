package Commands

import (
	"github.com/codecrafters-io/redis-starter-go/app/api"
	"net"
)

func PING(conn net.Conn) {
	api.PONG(conn)
}
