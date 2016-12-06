package main

import "fmt"

func plus(a int, b int) int {
	return a + b
}

func plus_plus(a, b, c int) int {
	return a + b + c
}

func main() {
	res := plus(1, 3)
	fmt.Println("1+3 = ", res)
	res = plus_plus(1, 2, 3)
	fmt.Println("1+2+3 = ", res)
}
