package tests

import (
	"github.com/dshulyak/borsh/tests/go-lib"
	"github.com/dshulyak/borsh/tests/imported"
)

type Alias [20]byte

type Local struct {
	LocalAlias  []Alias
	Bool        bool
	Uint32Slice []uint32
	Uint32Array [30]uint32
	PtrToSelf   *Local
}

type Hello struct {
	ID              imported.ByteAlias
	Ptr             *imported.Struct
	Ptrs            []*imported.Struct
	NestedPtrs      [][]*imported.Struct
	BytesIDs        []imported.ByteAlias
	BytesCollection imported.BytesCollection
	AnotherImport   *lib.Struct
}
