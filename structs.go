package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	fmt.Println(person{"Bob", 20})
	fmt.Println(person{name: "Alice", age: 22})
	p := person{name: "Sean", age: 32}
	fmt.Println(p)

	sp := &p

	fmt.Println(sp.age)
	sp.age = 52
	fmt.Println(sp.age)
}
