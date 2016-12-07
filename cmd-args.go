package main

import "os"
import "fmt"

func main() {

	args := os.Args
	only_args := os.Args[1:]

	arg := os.Args[3]
	fmt.Println(args)
	fmt.Println(only_args)
	fmt.Println(arg)
}
