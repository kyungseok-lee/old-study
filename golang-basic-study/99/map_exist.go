package main

func main() {

	m := map[int]string{
		1: "A",
		2: "A1",
		3: "A2",
		4: "A3",
	}

	val, exist := m[1]
	println(val, exist) // A true

	m2 := map[string]string{
		"apple":  "1",
		"banana": "2",
	}

	val2, exist2 := m2["apple"]
	println(val2, exist2) // 1 true

	val3, exist3 := m2["apple2"]
	println(val3, exist3) // false

	//if val3 == nil
	//	println()

	for key, val := range m2 {
		println("m2", key, val)
	}

}
