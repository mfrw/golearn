package main

import (
	"fmt"
	"time"
)

func f(from string) int {
	for i := 0; i < 50; i++ {
		fmt.Println(from, ":", i)
		time.Sleep(time.Millisecond * 10)
	}
	return 1
}

func main() {
	go f("direct")
	go f("goroutine")

	func(msg string) {
		fmt.Println(msg)
	}("going")

	var input string
	fmt.Scanln(&input)
	fmt.Println("done")
}
