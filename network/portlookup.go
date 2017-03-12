package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Fprintf(os.Stderr, "Usage: %s service type\n", os.Args[0])
		os.Exit(1)
	}
	nwType := os.Args[1]
	service := os.Args[2]

	port, err := net.LookupPort(nwType, service)
	if err != nil {
		fmt.Println("Error", err.Error())
		os.Exit(2)
	}
	fmt.Println("Service port ", port)
	os.Exit(0)
}
