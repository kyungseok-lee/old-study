package main

import "fmt"

func main() {
	msg := "Hello"
	fmt.Printf("01: %o\n", &msg)
	println("02: ", &msg)

	say(&msg)                   // 주소 전달
	fmt.Printf("03: %s\n", msg) // 변경된 메시지 출력
	fmt.Printf("03: %o\n", &msg) // 변경된 메시지 출력
}

func say(msg *string) {
	println(msg)
	println(&msg)
	println(*&msg)
	println(*msg)

	*msg = "Changed" // 메시지 변경
}
