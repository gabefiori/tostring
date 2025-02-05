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
	case int:
		return strconv.Itoa(t)
	case int8:
		return strconv.FormatInt(int64(t), 10)
	case int16:
		return strconv.FormatInt(int64(t), 10)
	case int32:
		return strconv.FormatInt(int64(t), 10)
	case int64:
		return strconv.FormatInt(t, 10)
	case uint:
		return strconv.FormatUint(uint64(t), 10)
	case uint8:
		return strconv.FormatUint(uint64(t), 10)
	case uint16:
		return strconv.FormatUint(uint64(t), 10)
	case uint32:
		return strconv.FormatUint(uint64(t), 10)
	case uint64:
		return strconv.FormatUint(t, 10)
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
	switch t := v.(type) {
	case []byte:
		return unsafe.String(unsafe.SliceData(t), len(t))
	default:
		return Any(v)
	}
}
