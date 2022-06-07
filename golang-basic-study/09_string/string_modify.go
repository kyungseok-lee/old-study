package main

import "fmt"

func main() {
	var s1 string = "hello, world"

	s1 = "abcdefg"

	fmt.Println(s1[0]) // 97
	fmt.Println('a')   // 97: ASCII 코드 a, 배열 형태로 각 문자에 접근

	s1[0] = 'z' // 컴파일 에러
}
