package borsh

import "io"

type Encodable interface {
	SizeBorsh() int
	MarshalBorsh(io.Writer) error
}

type Decodable interface {
	UnmarshalBorsh(io.Reader) error
}
