package main // special package for entry point

import (
	"fmt"
	"unsafe"
)

func main() { // entry point

	var _name1 string             // first method
	_new := "I really don't know" // second method
	var _name2 string = "smth"    // third method

	_name1 = "idk"
	// var (
	// 	firstName string = "Alex"
	// 	lastName string = "Skyba"
	// 	year int = 1000;
	// )
	fmt.Println("\n This is name 1 -> ", _name1,  " \n This is the name 2 -> ", _new,  " \n This is the name 3 -> ", _name2)
	fmt.Println("Size of int is ", unsafe.Sizeof(int(0)))

}
