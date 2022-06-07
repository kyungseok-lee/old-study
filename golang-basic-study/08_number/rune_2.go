package main

func main() {
	var r1 rune = "한"  // 컴파일 에러
	var r2 rune = '한글' // 컴파일 에러
	var r3 rune = "한글" // 컴파일 에러

	// # command-line-arguments
	// .\rune_2.go:4:6: cannot use "한" (type untyped string) as type rune in assignment
	// .\rune_2.go:5:16: invalid character literal (more than one character)
	// .\rune_2.go:6:6: cannot use "한글" (type untyped string) as type rune in assignment
}
