package main

import (
	"github.com/spacemeshos/borsh/gen"
	"github.com/spacemeshos/borsh/tests"
	"github.com/spacemeshos/borsh/tests/go-lib"
	"github.com/spacemeshos/borsh/tests/imported"
)

func main() {
	must(gen.Generate("imported", "tests/imported/types_borsh.go", imported.Struct{}))
	must(gen.Generate("lib", "tests/go-lib/types_borsh.go", lib.Struct{}))
	must(gen.Generate("tests", "tests/types_borsh.go", tests.Local{}, tests.Hello{}))
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
