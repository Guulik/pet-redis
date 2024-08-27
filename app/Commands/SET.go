package Commands

import (
	"fmt"
	"github.com/codecrafters-io/redis-starter-go/app/RESP"
	"github.com/codecrafters-io/redis-starter-go/app/Storage"
	"net"
)

func SET(conn net.Conn, key string, value string) {
	op := " SET "

	storage := Storage.GetInstance()
	if _, exists := storage.Get(key); exists {
		fmt.Println("op:", op, "The key already exists.")
	} else {
		storage.Set(key, value)
		fmt.Println("op:", op, "Key-Value pair set successfully.")
	}

	buf, err := RESP.EncodeSimpleString("OK")
	if err != nil {
		fmt.Println("op:", op, "failed to encode: ", err)
	}

	_, err = conn.Write(buf.Bytes())
	if err != nil {
		fmt.Println("op:", op, "failed to write response to client")
	}
}
