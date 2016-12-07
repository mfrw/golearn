package main

import (
	"fmt"
	"sort"
)

func main() {
	strs := []string{"c", "a", "b"}
	sort.Strings(strs)
	fmt.Println("Strings:", strs)

	ints := []int{7, 8, 9, 2, 1, 11, 5}
	sort.Ints(ints)
	fmt.Println("ints:", ints)
	fmt.Println("Sorted: ", sort.IntsAreSorted(ints))
}
