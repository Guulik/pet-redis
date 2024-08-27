package RESP

import (
	"bytes"
	"fmt"
	"github.com/tidwall/resp"
)

func EncodeSimpleString(s string) (bytes.Buffer, error) {
	var buf bytes.Buffer
	wr := resp.NewWriter(&buf)
	err := wr.WriteSimpleString(s)

	if err != nil {
		fmt.Println("failed to encode string with RESP")
		return bytes.Buffer{}, err
	}

	return buf, nil
}

func EncodeBulkString(s string) (bytes.Buffer, error) {
	var (
		buf bytes.Buffer
		err error
	)
	wr := resp.NewWriter(&buf)
	if s != "" {
		err = wr.WriteString(s)
	} else {
		err = wr.WriteNull()
	}

	if err != nil {
		fmt.Println("failed to encode string with RESP")
		return bytes.Buffer{}, err
	}

	return buf, nil
}
