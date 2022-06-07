package main

func main() {
	i := 0

L1:
	for {

		if i == 0 {
			break L1
		}
	}

	println("OK")
}
