package main

import (
	"github.com/spacemeshos/borsh/gen"
	"github.com/spacemeshos/borsh/testing"
)

func main() {
	if err := gen.Generate("testing", "./types_borsh.go", testing.Local{}, testing.Hello{}); err != nil {
		panic(err)
	}
}
