package main

import (
	"bytes"
	"fmt"
	"github.com/tidwall/resp"
	"io"
	"log"
	"net"
	"os"
	// Uncomment this block to pass the first stage
	// "net"
	// "os"
)

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	fmt.Println("Logs from your program will appear here!")

	l, err := net.Listen("tcp", "0.0.0.0:6379")
	if err != nil {
		fmt.Println("Failed to bind to port 6379")
		os.Exit(1)
	}
	con, err := l.Accept()
	if err != nil {
		fmt.Println("Error accepting connection: ", err.Error())
		os.Exit(1)
	}

	defer con.Close()

	buf := make([]byte, 1024)
	_, err = con.Read(buf)
	fmt.Println("read bytes from client: ", buf)

	HandleInput(con, buf)
}

func HandleInput(conn net.Conn, buffer []byte) {
	rd := resp.NewReader(bytes.NewReader(buffer))
	for {
		v, _, err := rd.ReadValue()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Read %s\n", v.Type())

		for i, v := range v.Array() {
			fmt.Printf("  #%d %s, value: '%s'\n", i, v.Type(), v)
		}
		if v.String() == "ping" {
			PINGResponse(conn)
		}
	}
}

func PINGResponse(conn net.Conn) {
	var buf bytes.Buffer
	wr := resp.NewWriter(&buf)
	err := wr.WriteSimpleString("PONG")
	if err != nil {
		fmt.Println("error occurred when encoding PONG with RESP")
	}

	_, err = conn.Write(buf.Bytes())
	if err != nil {
		fmt.Println("error occurred when writing PONG to client")
	}
}
