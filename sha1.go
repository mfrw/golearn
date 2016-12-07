package main

import (
	"crypto/sha1"
	"fmt"
)

func main() {
	s := "A test string for sha"
	h := sha1.New()

	h.Write([]byte(s))
	bs := h.Sum(nil)
	fmt.Println(s)
	fmt.Printf("%s\n", bs)
}
