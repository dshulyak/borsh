package testing

import (
	"github.com/spacemeshos/borsh/testing/imported"
)

type Local struct{}

type Hello struct {
	ID  imported.ByteAlias
	Ptr *imported.Struct
}
