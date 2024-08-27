package InputHandler

import (
	"bytes"
	"fmt"
	"github.com/codecrafters-io/redis-starter-go/app/Commands"
	"github.com/tidwall/resp"
	"io"
	"log"
	"net"
)

func Handle(conn net.Conn) {
	buf := make([]byte, 128)
	defer conn.Close()
	for {
		n, err := conn.Read(buf)
		if n == 0 {
			fmt.Println("No data to read.")
			break
		}
		if err != nil {
			fmt.Println("failed to read bytes")
		}
		rd := resp.NewReader(bytes.NewReader(buf[:n]))
		v, _, err := rd.ReadValue()
		fmt.Println("recieved bytes: ", buf, "readerValue: ", v)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		command := v.Array()[0].String()
		fmt.Println("command: ", command)

		switch command {
		case "PING":
			fmt.Println("ponging...")
			Commands.PINGResponse(conn)
		case "ECHO":
			fmt.Println("echoing...")
			phrase := v.Array()[1].String()
			Commands.ECHO(conn, phrase)
		case "SET":
			fmt.Println("setting...")
			key := v.Array()[1].String()
			value := v.Array()[2].String()
			Commands.SET(conn, key, value)
		case "GET":
			fmt.Println("getting...")
			key := v.Array()[1].String()
			Commands.GET(conn, key)
		default:
			fmt.Printf("Unknown command: %s\n", command)
		}
	}
}
