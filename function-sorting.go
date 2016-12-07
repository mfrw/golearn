package main

import (
	"fmt"
	"sort"
)

type bylen []string

func (s bylen) Len() int {
	return len(s)
}

func (s bylen) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s bylen) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}

func main() {
	fruits := []string{"peach", "banana", "apple", "mango", "kivi", "cherry"}
	sort.Sort(bylen(fruits))
	fmt.Println(fruits)
}
