package borsh

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"io"
)

func WriteLength(w io.Writer, lth int) error {
	return WriteUint32(w, uint32(lth))
}

func WriteUint8(w io.Writer, value uint8) error {
	return writeByte(w, value)
}

func WriteUint16(w io.Writer, value uint16) error {
	buf := make([]byte, 2)
	binary.LittleEndian.PutUint16(buf, value)
	_, err := write(w, buf)
	return err
}

func WriteUint64(w io.Writer, value uint64) error {
	buf := make([]byte, 8)
	binary.LittleEndian.PutUint64(buf, value)
	_, err := write(w, buf)
	return err
}

func WriteUint32(w io.Writer, value uint32) error {
	buf := make([]byte, 4)
	binary.LittleEndian.PutUint32(buf, value)
	_, err := write(w, buf)
	return err
}

func WriteBool(w io.Writer, val bool) error {
	if val {
		return writeByte(w, 1)
	}
	return writeByte(w, 0)
}

func WriteString(w io.Writer, val string) error {
	_, err := writeString(w, val)
	return err
}

func WriteBytes(w io.Writer, buf []byte) error {
	_, err := write(w, buf)
	return err
}

func write(w io.Writer, buf []byte) (int, error) {
	switch tw := w.(type) {
	case *bufio.Writer:
		return tw.Write(buf)
	case *bytes.Buffer:
		return tw.Write(buf)
	}
	return w.Write(buf)
}

func writeByte(w io.Writer, val byte) error {
	switch tw := w.(type) {
	case *bufio.Writer:
		return tw.WriteByte(val)
	case *bytes.Buffer:
		return tw.WriteByte(val)
	}
	_, err := w.Write([]byte{val})
	return err
}

func writeString(w io.Writer, val string) (int, error) {
	switch tw := w.(type) {
	case *bufio.Writer:
		return tw.WriteString(val)
	case *bytes.Buffer:
		return tw.WriteString(val)
	}
	return io.WriteString(w, val)
}
