package main

import "fmt"
import "time"

func main() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {
		for {
			j, more := <-jobs
			if more {
				fmt.Println("recieved job", j)
				time.Sleep(time.Second)
			} else {
				fmt.Println("recieved all jobs")
				done <- true
				return
			}
		}
	}()

	for j := 1; j <= 10; j++ {
		jobs <- j
		fmt.Println("sent job", j)
	}
	close(jobs)
	fmt.Println("sent all jobs")

	<-done
}
