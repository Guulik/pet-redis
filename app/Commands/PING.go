package Commands

import (
	"bytes"
	"fmt"
	"github.com/tidwall/resp"
	"net"
)

func PINGResponse(conn net.Conn) {
	var buf bytes.Buffer
	wr := resp.NewWriter(&buf)
	err := wr.WriteSimpleString("PONG")

	if err != nil {
		fmt.Println("failed to encode PONG with RESP")
	}

	_, err = conn.Write(buf.Bytes())
	if err != nil {
		fmt.Println("failed to write PONG to client")
	}
}
