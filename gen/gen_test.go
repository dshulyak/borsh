package gen

import (
	"testing"

	"github.com/stretchr/testify/require"
)

type Hello struct {
	Field1           uint32
	Field2           bool
	PtrToSelf        *Hello
	PtrToAnother     *Nested
	ByteSlice        []byte
	Uint32Slice      []uint32
	ByteArray        [20]byte
	StructSlice      []Nested
	PtrSlie          []*Nested
	NestedByteSlices [][][]byte
}

type Nested struct {
	Bytes []byte
}

func TestGenerate(t *testing.T) {
	require.NoError(t, GenerateFile("gen", "./hello_test.go", Nested{}, Hello{}))
}
