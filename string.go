// package tostring is a library for converting values of any type to strings when performance is critical.
package tostring

import (
	"fmt"
	"strconv"
	"unsafe"
)

// Any converts a value of any type to its string representation.
//
// It handles various types including string, fmt.Stringer, integers, booleans, and byte slices.
// For unsupported types, it uses fmt.Sprint to provide a default string representation.
func Any(v any) string {
	switch t := v.(type) {
	case string:
		return t
	case fmt.Stringer:
		return t.String()
	case int, int8, int16, int32, int64:
		return formatNum(t)
	case uint, uint8, uint16, uint32, uint64:
		return formatNum(t)
	case float32:
		return strconv.FormatFloat(float64(t), 'g', -1, 32)
	case float64:
		return strconv.FormatFloat(t, 'g', -1, 64)
	case []byte:
		return string(t)
	case nil:
		return "<nil>"
	case error:
		return t.Error()
	case bool:
		if t {
			return "true"
		}

		return "false"
	default:
		return fmt.Sprint(v)
	}
}

// AnyUnsafe is a faster alternative for converting byte slices to strings,
// but it uses an unsafe pointer, which can lead to undefined behavior if
// not used properly.
//
// For all other types, it delegates to the Any function for conversion.
func AnyUnsafe(v any) string {
	if b, ok := v.([]byte); ok {
		return unsafe.String(unsafe.SliceData(b), len(b))
	}

	return Any(v)
}

func formatNum(v any) string {
	switch t := v.(type) {
	case int:
		return itoa(t)
	case int8:
		return itoa(t)
	case int16:
		return itoa(t)
	case int32:
		return itoa(t)
	case int64:
		return itoa(t)
	case uint:
		return utoa(t)
	case uint8:
		return utoa(t)
	case uint16:
		return utoa(t)
	case uint32:
		return utoa(t)
	case uint64:
		return utoa(t)
	default:
		return ""
	}
}

func itoa[T int | int8 | int16 | int32 | int64](v T) string {
	return strconv.FormatInt(int64(v), 10)
}

func utoa[T uint | uint8 | uint16 | uint32 | uint64](v T) string {
	return strconv.FormatUint(uint64(v), 10)
}
