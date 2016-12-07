package main

import "fmt"

func intseq(initial int) func() int {
	i := initial
	return func() int {
		i += 1
		return i
	}
}

func jumperseq(initial int) func(a int) int {
	i := initial
	return func(a int) int {
		i += a
		return i
	}
}

func main() {
	nextint := intseq(10)
	fmt.Println(nextint())
	fmt.Println(nextint())
	jum := jumperseq(6)
	fmt.Println(jum(2))

}
