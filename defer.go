package main

import (
	"fmt"
	"os"
)

func main() {

	f := create_file("/tmp/defer.txt")
	defer close_file(f)
	write_file(f)
}

func create_file(p string) *os.File {
	fmt.Println("creating")
	f, err := os.Create(p)
	if err != nil {
		panic(err)
	}
	return f
}

func write_file(f *os.File) {
	fmt.Println("Writing")
	fmt.Fprintln(f, "data")
}

func close_file(f *os.File) {
	fmt.Println("Closing")
	f.Close()
}
