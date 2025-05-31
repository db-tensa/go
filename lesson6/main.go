package main

import (
	"fmt"
)

type Students struct {
	Name     string
	Age    int
}

func main() {

	a := Students{Name: "ann", Age: 25}
	b := &a
	b.Name = "Bob"
	fmt.Println(a.Name)
	fmt.Println(b.Name)
}
