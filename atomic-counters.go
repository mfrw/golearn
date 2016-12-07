package main

import "fmt"
import "time"
import "sync/atomic"
import "runtime"

func main() {
	var ops uint64 = 0

	for i := 0; i < 50000; i++ {
		go func() {
			for {
				atomic.AddUint64(&ops, 1)
				runtime.Gosched()
			}
		}()
	}
	time.Sleep(time.Second)

	ops_final := atomic.LoadUint64(&ops)
	fmt.Println("ops:", ops_final)
}
