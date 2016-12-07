package main

import (
	"bytes"
	"fmt"
	"regexp"
)

func main() {
	match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
	fmt.Println(match)

	r, _ := regexp.Compile("p([a-z]+)ch")

	fmt.Println(r.MatchString("peach"))

	fmt.Println(r.FindString("peach"))
	fmt.Println(r.FindAllString("peach punch pinch", -1))

	fmt.Println(r.Match([]byte("peach")))
	r = regexp.MustCompile("p([a-z]+)ch")
	fmt.Println(r)

	in := []byte("a peach")
	out := r.ReplaceAllFunc(in, bytes.ToUpper)
	fmt.Println(string(out))
}
