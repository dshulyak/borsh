package main

import (
	"github.com/spacemeshos/borsh/gen"
	"github.com/spacemeshos/borsh/testing"
	"github.com/spacemeshos/borsh/testing/imported"
)

func main() {
	must(gen.Generate("testing", "testing/imported/types_borsh", imported.Struct{}))
	must(gen.Generate("testing", "testing/types_borsh.go", testing.Hello{}))
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
