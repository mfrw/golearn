package main

import (
	"encoding/json"
	"fmt"
	//	"os"
)

type resp struct {
	page   int
	fruits []string
}

type resp2 struct {
	page   int      `json:"page"`
	fruits []string `json:"fruits"`
}

func main() {
	bolb, _ := json.Marshal(true)
	fmt.Println(string(bolb))

	intb, _ := json.Marshal(1)
	fmt.Println(string(intb))

	fltb, _ := json.Marshal(3.14)
	fmt.Println(string(fltb))

}
