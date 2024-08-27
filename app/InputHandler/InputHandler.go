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
		if err != nil {
			fmt.Println("failed to read bytes")
		}
		if n == 0 {
			fmt.Println("No more data to read.")
			break
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

		fmt.Println("v first: ", v.Array()[0])

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
		default:
			fmt.Printf("Unknown command: %s\n", command)
		}

		/*for i, v := range v.Array() {
			fmt.Printf("  #%d %s, value: '%s'\n", i, v.Type(), v)
			fmt.Println("v string: ", v.String())

			if strings.EqualFold(v.String(), "ping") {
				fmt.Println("ponging...")
				Commands.PINGResponse(conn)
				break
			}

		}*/
	}
}
