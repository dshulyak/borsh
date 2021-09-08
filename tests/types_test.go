package tests_test

import (
	"bytes"
	"math/rand"
	"reflect"
	"testing"
	"testing/quick"

	fuzz "github.com/google/gofuzz"
	"github.com/spacemeshos/borsh/tests"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFuzzDecode(t *testing.T) {
	assert.NoError(t, quick.Check(func(h1 *tests.Hello) bool {
		var buf bytes.Buffer
		if !assert.NoError(t, h1.MarshalBorsh(&buf)) {
			return false
		}
		var h2 tests.Hello
		if !assert.NoError(t, h2.UnmarshalBorsh(&buf)) {
			return false
		}
		return assert.Equal(t, h1, &h2)
	}, &quick.Config{
		MaxCountScale: 100,
		Values: func(values []reflect.Value, rng *rand.Rand) {
			require.Len(t, values, 1)
			fuzzer := fuzz.New()
			fuzzer = fuzzer.RandSource(rng)
			var h1 tests.Hello
			fuzzer.Fuzz(&h1)
			values[0] = reflect.ValueOf(&h1)
		},
	}))
}

func TestFuzzNoCrash(t *testing.T) {
	assert.NoError(t, quick.Check(func(buf []byte) bool {
		b := bytes.NewBuffer(buf)
		return assert.NotPanics(t, func() {
			var h tests.Hello
			h.UnmarshalBorsh(b)
		})
	}, &quick.Config{
		MaxCountScale: 1000,
		Values: func(values []reflect.Value, rng *rand.Rand) {
			require.Len(t, values, 1)
			buf := []byte{}
			fuzzer := fuzz.New()
			fuzzer = fuzzer.RandSource(rng)
			fuzzer.Fuzz(&buf)
			values[0] = reflect.ValueOf(buf)
		},
	}))
}

func TestRecursive(t *testing.T) {
	l1 := tests.Local{PtrToSelf: &tests.Local{Bool: true}}
	var b bytes.Buffer
	require.NoError(t, l1.MarshalBorsh(&b))
	var l2 tests.Local
	require.NoError(t, l2.UnmarshalBorsh(&b))
	require.Equal(t, l1, l2)
}
