package main

import (
	"fmt"
)

type Students struct {
	Name     string
	Age    int
}

func main() {

	var x int  = 42
	var p *int = &x 
	fmt.Println(p)

}
