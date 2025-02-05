# tostring
![test badge](https://github.com/gabefiori/tostring/actions/workflows/go.yml/badge.svg) [![Go Report Card](https://goreportcard.com/badge/github.com/gabefiori/tostring)](https://goreportcard.com/report/github.com/gabefiori/tostring) [![PkgGoDev](https://pkg.go.dev/badge/github.com/gabefiori/tostring)](https://pkg.go.dev/github.com/gabefiori/tostring)

A library for converting values of any type to strings when performance is critical.

It is designed to be faster than the standard `fmt` package by using specialized conversions for supported types.

## Support
- `string`
- `fmt.Stringer`
- `bool`
- `[]byte`
- `error`
- Integers (`int`, `int8`, `int16`, `int32`, `int64`)
- Unsigned integers (`uint`, `uint8`, `uint16`, `uint32`, `uint64`)
- Floats (`float32`, `float64`)
- Handles `nil` values gracefully.

## Installation
To use the `tostring` package, you can install it using:

```sh
go get github.com/gabefiori/tostring
```

## Usage
### Any 
The `Any` function converts a value of any type to its string representation.

```go
package main

import (
    "fmt"
    "github.com/gabefiori/tostring"
)

func main() {
    tostring.Any("Hello, World!")  // -> "Hello, World!"
    tostring.Any(42)               // -> "42"
    tostring.Any(true)             // -> "true"
    tostring.Any([]byte("hello"))  // -> "hello"
    tostring.Any(nil)              // -> "<nil>"
}
```

### AnyUnsafe 
The `AnyUnsafe` function provides a faster way to convert byte slices to strings using unsafe pointers.
Use this function with caution, as improper use can lead to undefined behavior.

For all other types, it delegates to the `Any` function for conversion.
```go
package main

import (
    "fmt"
    "github.com/gabefiori/tostring"
)

func main() {
    tostring.AnyUnsafe("Hello, World!")  // -> "Hello, World!"
    tostring.AnyUnsafe(42)               // -> "42"
    tostring.AnyUnsafe(true)             // -> "true"
    tostring.AnyUnsafe([]byte("hello"))  // -> "hello"
    tostring.AnyUnsafe(nil)              // -> "<nil>"
}
```

## Benchmarks
To run the benchmarks, use the following command:
```sh
go test ./... -bench=. -benchmem
```
