package tests

import (
	"github.com/spacemeshos/borsh/tests/go-lib"
	"github.com/spacemeshos/borsh/tests/imported"
)

type Local struct {
	Bool        bool
	Uint32Slice []uint32
	Uint32Array [30]uint32
	PtrToSelf   *Local
	PtrToHello  *Hello
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
