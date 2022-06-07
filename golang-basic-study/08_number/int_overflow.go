package main

import (
	"math"
)

func main() {
	var num1 uint8 = math.MaxUint8 + 1   // 오버플로우 컴파일 에러 발생
	var num2 uint16 = math.MaxUint16 + 1 // 오버플로우 컴파일 에러 발생
	var num3 uint32 = math.MaxUint32 + 1 // 오버플로우 컴파일 에러 발생
	var num4 uint64 = math.MaxUint64 + 1 // 오버플로우 컴파일 에러 발생

	// # command-line-arguments
	// .\int_overflow.go:8:6: constant 256 overflows uint8
	// .\int_overflow.go:9:6: constant 65536 overflows uint16
	// .\int_overflow.go:10:6: constant 4294967296 overflows uint32
	// .\int_overflow.go:11:6: constant 18446744073709551616 overflows uint64
}
