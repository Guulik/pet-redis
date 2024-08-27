package Commands

import (
	"bytes"
	"fmt"
	"github.com/tidwall/resp"
	"net"
)

func ECHO(conn net.Conn, phrase string) {
	var buf bytes.Buffer
	wr := resp.NewWriter(&buf)
	err := wr.WriteSimpleString(phrase)
	if err != nil {
		fmt.Println("failed to encode phrase with RESP")
	}

	_, err = conn.Write(buf.Bytes())
	if err != nil {
		fmt.Println("failed to write response to client")
	}
}
