package InputHandler

import (
	"bytes"
	"fmt"
	"github.com/codecrafters-io/redis-starter-go/app/Commands"
	"github.com/tidwall/resp"
	"io"
	"log"
	"net"
	"strings"
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

		for i, v := range v.Array() {
			fmt.Printf("  #%d %s, value: '%s'\n", i, v.Type(), v)

			if strings.EqualFold(v.String(), "ping") {
				fmt.Println("ponging...")
				Commands.PINGResponse(conn)
				break
			}

		}
	}
}
