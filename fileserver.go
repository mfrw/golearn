package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// A static web server serving a path
	fmt.Println("Server starting on Port 8080")
	log.Fatal(http.ListenAndServe(":8080", http.FileServer(http.Dir("/usr/share/doc"))))
}
