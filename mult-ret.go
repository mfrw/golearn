package main

import "fmt"

func vals(a, b, c int) (int, int) {
	return a + b + c, a + c
}

func main() {
	a, b := vals(1, 2, 3)
	fmt.Println(a)
	fmt.Println(b)

	_, c := vals(2, 3, 4)
	fmt.Println(c)
}
