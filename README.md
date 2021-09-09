Codegen for borsh codec
===

Follows specification defined in https://borsh.io/.

How to use?
===

Define an executable with a sample content:

```go
import (
	"github.com/dshulyak/borsh/gen"
)

func main() {
	must(gen.Generate("package", "types_borsh.go", Struct{}))
}
```

It will generate following methods for `Struct` type.

```go
type Encodable interface {
	SizeBorsh() int
	MarshalBorsh(io.Writer) error
}

type Decodable interface {
	UnmarshalBorsh(io.Reader) error
}
```

Considerations
===

Nil ptr and nil byte slices are encoded as an Option. 