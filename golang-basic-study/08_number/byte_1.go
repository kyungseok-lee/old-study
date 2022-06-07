package main

import "fmt"

func main() {
	var num1 byte = 10
	var num2 byte = 0x32
	var num3 byte = 'a'
	fmt.Println(num1, num2, num3) // 10 50 97

	var num4 byte = '!'
	fmt.Println(num4) // 33
}
