package main

func main() {

	s := make([]int, 5, 10)

	for i := 0; i < 50; i++ {
		s = append(s, i)
		println(len(s), cap(s), i) // cap 10, 20, 40, 80
	}

}
