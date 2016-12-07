package main

import "time"
import "fmt"

func main() {
	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)

	limiter := time.Tick(time.Millisecond * 200)

	for req := range requests {
		<-limiter
		fmt.Println("requests", req, time.Now())
	}
	bursty_limiter := make(chan time.Time, 3)

	for i := 0; i < 3; i++ {
		bursty_limiter <- time.Now()
	}

	go func() {
		for t := range time.Tick(time.Millisecond * 200) {
			bursty_limiter <- t
		}
	}()

	bursty_requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		bursty_requests <- i
	}
	close(bursty_requests)
	for req := range bursty_requests {
		<-bursty_limiter
		fmt.Println("request", req, time.Now())
	}
}
