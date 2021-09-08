package borsh

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"io"
)

func ReadBool(r io.Reader) (bool, error) {
	byt, err := readByte(r)
	if err != nil {
		return false, err
	}
	return byt == 1, nil
}

func ReadBytes(r io.Reader, buf []byte) error {
	_, err := io.ReadFull(r, buf)
	return err
}

func ReadUint32(r io.Reader) (uint32, error) {
	buf := make([]byte, 4)
	_, err := io.ReadFull(r, buf)
	if err != nil {
		return 0, err
	}
	return binary.LittleEndian.Uint32(buf), nil
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
