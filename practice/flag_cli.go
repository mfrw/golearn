package main

import (
	"flag"
	"fmt"
)

var name = flag.String("name", "World", "A name to say hello to.")
var spanish bool

func init() {
	flag.BoolVar(&spanish, "spanish", false, "Use spanish lang.")
	flag.BoolVar(&spanish, "s", false, "Use spanish lang.")
}

func main() {
	flag.Parse()
	if spanish == true {
		fmt.Println("Holla %s!", *name)
	} else {
		fmt.Println("Hello %s!", *name)
	}
}
