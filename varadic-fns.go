package main

import "fmt"

func sums(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

func main() {
	fmt.Println(sums(1, 2, 3, 4))
	fmt.Println(sums(1, 2))

	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	sums(nums...)
}
