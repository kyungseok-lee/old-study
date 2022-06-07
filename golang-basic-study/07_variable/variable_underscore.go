package main

import "fmt"

func main() {
	a := 1
	b := "hello"

	_ = b // 미사용 변수로 인한 컴파일 에러 방지

	fmt.Println(a)
}
