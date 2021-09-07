package borsh

import (
	"bufio"
	"bytes"
	"io"
)

func ReedBool(r io.Reader) (bool, error) {
	return false, nil
}

func readByte(r io.Reader) (byte, error) {
	switch r := r.(type) {
	case *bytes.Buffer:
		return r.ReadByte()
	case *bytes.Reader:
		return r.ReadByte()
	case *bufio.Reader:
		return r.ReadByte()
	case io.ByteReader:
		return r.ReadByte()
	}
	buf := make([]byte, 1)
	_, err := io.ReadFull(r, buf)
	return buf[0], err
}
