package main

import (
	"github.com/spacemeshos/borsh/gen"
	"github.com/spacemeshos/borsh/testing/imported"
)

type Hello struct {
	ID imported.ByteAlias
}

func main() {
	gen.GenerateFile("main", "./hello.go", Hello{})
}
