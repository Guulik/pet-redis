package Commands

import (
	"fmt"
	"github.com/codecrafters-io/redis-starter-go/app/RESP"
	"github.com/codecrafters-io/redis-starter-go/app/Storage"
	"net"
)

func GET(conn net.Conn, key string) {
	op := " SET "

	storage := Storage.GetInstance()
	if value, exists := storage.Get(key); exists {

		buf, err := RESP.EncodeBulkString(value)
		if err != nil {
			fmt.Println("op:", op, "failed to encode: ", err)
		}

		_, err = conn.Write(buf.Bytes())
		if err != nil {
			fmt.Println("op:", op, "failed to write response to client")
		}
	}

}
