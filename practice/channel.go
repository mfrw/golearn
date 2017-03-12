package main

import (
	"fmt"
	"time"
)

func printCount(c chan int) {
	//defer(fmt.Println("Exiting from printCount"))
	num := 0
	for num >= 0 {
		num = <-c
		fmt.Println(num, " ")
	}
}

func main() {
	c := make(chan int)
	a := []int{8, 6, 7, 5, 3, 0, 9, -1}

	go printCount(c)

	for _, v := range a {
		c <- v
		time.Sleep(time.Millisecond * 100)
	}
	time.Sleep(time.Millisecond * 100)
	fmt.Println("End of main")
}
