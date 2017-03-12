package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	resp, _ := http.Get("http://google.com")
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("body", string(body))
	resp.Body.Close()
}
