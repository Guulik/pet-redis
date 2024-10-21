package api

import (
	"fmt"
	"github.com/codecrafters-io/redis-starter-go/app/RESP"
	"net"
)

func PONG(conn net.Conn) {
	buf, err := RESP.EncodeSimpleString("PONG")
	if err != nil {
		fmt.Println("failed to encode:", err)
	}

	_, err = conn.Write(buf.Bytes())
	if err != nil {
		fmt.Println("failed to write PONG to client")
	}
}

func OK(conn net.Conn) {
	op := "OK response"

	buf, err := RESP.EncodeSimpleString("OK")
	if err != nil {
		fmt.Println("op:", op, "failed to encode: ", err)
	}

	_, err = conn.Write(buf.Bytes())
	if err != nil {
		fmt.Println("op:", op, "failed to write response to client")
	}
}
