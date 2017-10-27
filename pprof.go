package main

// usage : ./trace 2>trace.out 
// go tool trace trace.out

import (
        "fmt"
        "os"
        "runtime/trace"
        "sync"
)

var wg sync.WaitGroup

func work() {
        wg.Add(10)
        for i := 0; i < 10; i++ {
                go func() {
                        fib(42)
                        wg.Done()

                }()
        }
        wg.Wait()
}

func fib(n int) int {
        if n < 2 {
                return n
        }
        return fib(n-1) + fib(n-2)
}

// go: noinline
func expensive() string {
        a := ""
        for i := 0; i < 1000; i++ {
                a += fmt.Sprintf("This is a string %d", i)
                a = a + "some more shit"
        }

        return a
}

func main() {

        pprof.StartCPUProfile(os.Stderr)
        defer pprof.Stop()

        for i := 0; i < 5; i++ {
                work()
                b := expensive()
                _ = b
        }

}
