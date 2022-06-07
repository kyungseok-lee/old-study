package main

func main() {
	var num1 uint8 = 0 - 1  // 오버플로우 컴파일 에러 발생
	var num2 uint16 = 0 - 1 // 오버플로우 컴파일 에러 발생
	var num3 uint32 = 0 - 1 // 오버플로우 컴파일 에러 발생
	var num4 uint64 = 0 - 1 // 오버플로우 컴파일 에러 발생

	// # command-line-arguments
	// .\int_underflow.go:4:6: constant -1 overflows uint8
	// .\int_underflow.go:5:6: constant -1 overflows uint16
	// .\int_underflow.go:6:6: constant -1 overflows uint32
	// .\int_underflow.go:7:6: constant -1 overflows uint64
}
