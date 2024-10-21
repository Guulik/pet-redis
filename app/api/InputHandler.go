package api

import (
	"bytes"
	"fmt"
	"github.com/codecrafters-io/redis-starter-go/app/Commands"
	"github.com/tidwall/resp"
	"io"
	"log"
	"net"
	"strings"
	"time"
)

func Execute(conn *net.Conn, command string, args []resp.Value) {
	switch command {
	case "PING":
		fmt.Println("ponging...")
		Commands.PING(*conn)
	case "ECHO":
		fmt.Println("echoing...")
		phrase := args[1].String()
		Commands.ECHO(*conn, phrase)
	case "SET":
		fmt.Println("setting...")
		key := args[1].String()
		value := args[2].String()
		ttlType := args[3].String()
		expire := args[4].String()

		var (
			ttl time.Duration
			err error
		)
		if strings.EqualFold(ttlType, "px") {
			ttl, err = time.ParseDuration(expire + "ms")
			if err != nil {
				fmt.Println("failed to parse duration")
			}
			Commands.SET(*conn, key, value, Commands.WithTTL(ttl))
			break
		}
		Commands.SET(*conn, key, value)
	case "GET":
		fmt.Println("getting...")
		key := args[1].String()
		Commands.GET(*conn, key)
	default:
		fmt.Printf("Unknown command: %s\n", command)
	}
}

func ReadInput(conn net.Conn) {
	buf := make([]byte, 128)
	defer conn.Close()
	//why infinite FOR loop????
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

		args := v.Array()[1:]
		Execute(&conn, command, args)
	}
}
