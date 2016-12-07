package main

import "os"

func main() {
	panic("Something bad happened")

	_, err := os.Create("/tmp/file")
	if err != nil {
		panic(err)
	}
}
