package main

func main() {
	//ex01()
	//ex02()
	//ex03()
	//ex04()
	ex05()
}

func ex01() {
	s := make([]int, 5, 10) // (타입, Length(슬라이스의 길이), Capacity(내부 배열의 최대 길이))

	println(len(s), cap(s)) // len 5, cap 10
}

func ex02() {
	var s []int

	if s == nil {
		println("Nil Slice")
	}

	println(len(s), cap(s)) // 길이, 용량
}

func ex03() {
	var s []string

	if s == nil {
		println("Nil Slice")
	}

	println(len(s), cap(s))
}

func ex04() {
	a := []string{"1", "2", "3"}

	println("ex04")
	println(&a)
	println(a)
	println(a[0])
	println(a[1])
	println(a[2])
	// println(a[3])
}

func ex05() {
	a := make([]string, 3, 5)

	//for index, name := range names {
	//	println(index, name)
	//}

	a[0] = "0a"
	a[1] = "1a"
	a[2] = "2a"
	//a[3] = "3a" // error

	for index, value := range a {
		println("for ", index, value)
	}

}
