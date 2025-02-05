package tostring

import (
	"fmt"
	"strings"
	"testing"
)

func BenchmarkOther(b *testing.B) {
	for _, t := range others {
		b.Run(t.name, func(b *testing.B) {
			b.Run("Any", func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					_ = Any(t.value)
				}
			})

			b.Run("Sprint", func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					_ = fmt.Sprint(t.value)
				}
			})
		})
	}
}

func BenchmarkByteSlice(b *testing.B) {
	bb := []byte(strings.Repeat("val", 5))

	b.Run("Safe", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = Any(bb)
		}
	})

	b.Run("Unsafe", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = AnyUnsafe(bb)
		}
	})
}

var smallDigits = []test{
	{"Int", int(10)},
	{"Int8", int8(10)},
	{"Int16", int16(10)},
	{"Int32", int32(10)},
	{"Int64", int64(10)},
	{"Uint", uint(10)},
	{"Uint8", uint8(10)},
	{"Uint16", uint16(10)},
	{"Uint32", uint32(10)},
	{"Uint64", uint64(10)},
	{"Float32", float32(10)},
	{"Float64", float64(10)},
}

func BenchmarkIntegersSmall(b *testing.B) {
	for _, t := range smallDigits {
		b.Run(t.name, func(b *testing.B) {
			b.Run("Any", func(b *testing.B) {
				b.ReportAllocs()
				for i := 0; i < b.N; i++ {
					_ = Any(t.value)
				}
			})

			b.Run("Sprint", func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					_ = fmt.Sprint(t.value)
				}
			})
		})
	}
}

func BenchmarkIntegers(b *testing.B) {
	for _, t := range integers {
		b.Run(t.name, func(b *testing.B) {
			b.Run("Any", func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					_ = Any(t.value)
				}
			})

			b.Run("Sprint", func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					_ = fmt.Sprint(t.value)
				}
			})
		})
	}
}
