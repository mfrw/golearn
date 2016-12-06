package main

import "fmt"

func main() {
	s := make([]string, 3)
	fmt.Println("emp: ", s)
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"

	fmt.Println("len:", len(s))

	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apended:", s)
}
