package Commands

import (
	"fmt"
	"github.com/codecrafters-io/redis-starter-go/app/RESP"
	"net"
)

func PINGResponse(conn net.Conn) {
	buf, err := RESP.EncodeSimpleString("PONG")
	if err != nil {
		fmt.Println("failed to encode:", err)
	}

	_, err = conn.Write(buf.Bytes())
	if err != nil {
		fmt.Println("failed to write PONG to client")
	}
}
