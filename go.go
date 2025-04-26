package main // special package for entry point

import (
	"fmt"
	"unsafe"
)

func main() { // entry point

	var _name1 string             // first method
	_new := "I really don't know" // second method
	var _name2 string = "smth"    // third method


	// if doing variable + := data 
	//then you cannot set a type of variable even a const 

	_name1 = "idk"

	const PI = 3.14;
	 const smth = "smth"
	if true {
		value := 20
		fmt.Println(value)
	}
	// var (
	// 	firstName string = "Alex"
	// 	lastName string = "Skyba"
	// 	year int = 1000;
	// )
	fmt.Println("\n This is name 1 -> ", _name1,  " \n This is the name 2 -> ", _new,  " \n This is the name 3 -> ", _name2)

	fmt.Println("\n Size of int is ", unsafe.Sizeof(int(0)))
	fmt.Println("\n This is PI (const) -> ", PI, "\n This is smth -> (const) ", smth)

}
